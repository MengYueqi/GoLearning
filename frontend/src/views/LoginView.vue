<template>
  <div class="login-container">
    <h2>Login</h2>
    <form @submit.prevent="handleLogin">
      <div class="form-group">
        <label for="username">Username</label>
        <input type="text" v-model="username" id="username" required />
      </div>
      <div class="form-group">
        <label for="password">Password</label>
        <input type="password" v-model="password" id="password" required />
      </div>
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';
import router from "@/router";
import {jwtDecode} from "jwt-decode";

export default {
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    async handleLogin() {
      try {
        const response = await axios.post('http://localhost:8081/api/login', {
          username: this.username,
          password: this.password
        });
        console.log('Login successful:', response.data);

        // 存储 token 在 localStorage
        const token = response.data.token;
        localStorage.setItem('token', token);

        // 解码 JWT 获取用户信息
        const decoded = jwtDecode(token);
        console.log('Decoded JWT:', decoded);

        // 跳转到首页
        router.push('/allBlogs');
      } catch (error) {
        console.error('Error during login:', error.response ? error.response.data : error.message);
        // 在这里处理登录失败的情况
      }
    }
  }
};
</script>

<style scoped>
.login-container {
  max-width: 300px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
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
</style>
