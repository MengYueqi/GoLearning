<template>
  <div>
    <h2>Top {{ N }} Liked Blogs</h2>
    <ul v-if="blogs.length > 0">
      <li v-for="blog in blogs" :key="blog.Id">
        <h3>Blog ID: {{ blog.Id }}</h3>
        <p>{{ blog.Content }}</p>
        <p><strong>Likes:</strong> {{ blog.Likes }}</p>
        <small>Created At: {{ formatDate(blog.CreatedAt) }}</small>
      </li>
    </ul>
    <p v-else>Loading or no data available.</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

// 显示前几名
const N = 3
const blogs = ref([])

const fetchTopBlogs = async () => {
  try {
    const response = await axios.post('http://localhost:8081/api/getNTopBlogs', {
      N: N
    })
    blogs.value = response.data.blogs || []
  } catch (err) {
    console.error('Failed to fetch top blogs:', err)
  }
}

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  return date.toLocaleString()
}

onMounted(fetchTopBlogs)
</script>

<style scoped>
h2 {
  margin-bottom: 20px;
}
li {
  padding: 15px;
  border: 1px solid #ccc;
  margin-bottom: 10px;
  border-radius: 5px;
}
p {
  margin: 5px 0;
}
small {
  color: gray;
}
</style>
