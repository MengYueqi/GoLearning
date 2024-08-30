import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/LoginView.vue')
    },
    {
        path: '/home',
        name: 'Home',
        component: () => import('@/views/HomePageView.vue')
    },
    {
        path: '/blogs',
        name: 'Blogs',
        component: () => import('@/views/BlogList.vue')
    },
    {
        path: '/addBlog',
        name: 'addBlog',
        component: () => import('@/views/AddBlog.vue')
    },
    {
        path: '/allBlogs',
        name: 'allBlogs',
        component: () => import('@/views/AllBlogsList.vue')
    }
]


const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router;