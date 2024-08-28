<template>
  <div class="home-page">
    <h1>Welcome to Home Page</h1>
    <p>This is the home page of your application. You can customize this page to show your content.</p>

    <button @click="handleButtonClick">Click Me</button>
  </div>
</template>

<script>
export default {
  name: 'HomePageView',
  data() {
    return {
      // 可以在这里定义组件的状态数据
    };
  },
  methods: {
    handleButtonClick() {
      // 点击按钮后的逻辑
      alert('Button clicked!');
    }
  }
};
</script>

<style scoped>
.home-page {
  padding: 20px;
  text-align: center;
}

h1 {
  color: #333;
  margin-bottom: 20px;
}

p {
  font-size: 16px;
  color: #666;
}

button {
  padding: 10px 20px;
  font-size: 16px;
  color: white;
  background-color: #007bff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
