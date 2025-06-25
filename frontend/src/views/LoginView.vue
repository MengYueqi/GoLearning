<template>
  <div class="login-container">
    <h2>{{ isLogin ? 'Login' : 'Register' }}</h2>
    <form @submit.prevent="isLogin ? handleLogin() : handleRegister()">
      <div class="form-group">
        <label for="username">Username</label>
        <input type="text" v-model="username" id="username" required />
      </div>
      <div v-if="!isLogin" class="form-group">
        <label for="email">Email</label>
        <input type="email" v-model="email" id="email" required />
      </div>
      <div class="form-group">
        <label for="password">Password</label>
        <input type="password" v-model="password" id="password" required />
      </div>
      <div v-if="!isLogin" class="form-group">
        <label for="confirmPassword">Confirm Password</label>
        <input type="password" v-model="confirmPassword" id="confirmPassword" required />
      </div>
      <button type="submit">{{ isLogin ? 'Login' : 'Register' }}</button>
    </form>
    <p @click="toggleForm" class="toggle-link">
      {{ isLogin ? "Don't have an account? Register" : "Already have an account? Login" }}
    </p>
  </div>
</template>

<script>
import axios from 'axios';
import router from '@/router';
import { jwtDecode } from 'jwt-decode';

export default {
  data() {
    return {
      isLogin: true,
      username: '',
      email: '',
      password: '',
      confirmPassword: ''
    };
  },
  methods: {
    toggleForm() {
      this.isLogin = !this.isLogin;
      this.username = '';
      this.email = '';
      this.password = '';
      this.confirmPassword = '';
    },
    async handleLogin() {
      try {
        const response = await axios.post('http://localhost:8081/api/login', {
          username: this.username,
          password: this.password
        });

        const token = response.data.token;
        localStorage.setItem('token', token);

        const decoded = jwtDecode(token);
        console.log('Decoded JWT:', decoded);

        router.push('/allBlogs');
      } catch (error) {
        console.error('Login failed:', error.response?.data || error.message);
        alert('Login failed: ' + (error.response?.data?.message || error.message));
      }
    },
    async handleRegister() {
      if (this.password !== this.confirmPassword) {
        alert('Passwords do not match.');
        return;
      }
      try {
        const response = await axios.post('http://localhost:8081/api/register', {
          username: this.username,
          email: this.email,
          password: this.password
        });

        // 后端可能返回的是 error 字段
        if (response.data && response.data.error) {
          alert('Registration failed: ' + response.data.error);
        } else {
          alert('Registration successful! Please log in.');
          this.toggleForm(); // 切换到登录表单
        }
      } catch (error) {
        // 处理请求失败的情况，如 500、网络错误
        if (error.response && error.response.data && error.response.data.error) {
          alert('Registration failed: ' + error.response.data.error);
        } else {
          alert('Registration failed: ' + error.message);
        }
      }
    }
  }
};
</script>

<style scoped>
.login-container {
  max-width: 320px;
  margin: 0 auto;
  padding: 25px;
  border: 1px solid #ccc;
  border-radius: 8px;
}

h2 {
  text-align: center;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 6px;
}

input {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}

.toggle-link {
  margin-top: 15px;
  text-align: center;
  color: #007bff;
  cursor: pointer;
  text-decoration: underline;
}
</style>
