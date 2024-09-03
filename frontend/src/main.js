import { createApp } from 'vue'
import App from './App.vue'
import router from './router';
import axios from "axios";

// 在 Axios 请求拦截器中全局设置 JWT 令牌
axios.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token');  // 获取 JWT 令牌
        console.log(token)
        if (token) {
            config.headers.Authorization = `${token}`;  // 设置请求头
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);


createApp(App).use(router).mount('#app');
