<template>
  <div>
    <h2>Blogs by Author {{ authorId }}</h2>
    <ul v-if="blogs.length > 0">
      <li v-for="blog in blogs" :key="blog.Id">
        <h3>{{ blog.Username }} (ID: {{ blog.Id }})</h3>
        <p>{{ blog.Content }}</p>
        <small>Created At: {{ formatDate(blog.CreatedAt) }}</small>
        <button @click="deleteBlog(blog.Id)">Delete</button>
      </li>
    </ul>
    <p v-else>Loading blogs or no blogs available.</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const authorId = ref(null)
const blogs = ref([])

const fetchBlogs = async () => {
  try {
    const response = await axios.post('http://localhost:8081/GetAllBlogsById', {
      userId: 1
    })
    const data = response.data
    authorId.value = data.authorId
    blogs.value = data.blogs
  } catch (error) {
    console.error('Error fetching blogs:', error)
  }
}

const deleteBlog = async (blogId) => {
  try {
    const response = await axios.post('http://localhost:8081/deleteBlog', {
      blogId: blogId
    })
    if (response.data.status === 'success') {
      // 删除成功后，从博客列表中移除该博客
      blogs.value = blogs.value.filter(blog => blog.Id !== blogId)
    } else {
      console.error('Failed to delete blog.')
    }
  } catch (error) {
    console.error('Error deleting blog:', error)
  }
}

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleString()
}

onMounted(fetchBlogs)
</script>

<style scoped>
h2 {
  margin-bottom: 20px;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  margin-bottom: 20px;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
  position: relative;
}
h3 {
  margin: 0 0 5px;
}
p {
  margin: 0 0 10px;
}
small {
  color: #666;
}
button {
  padding: 5px 10px;
  border: none;
  border-radius: 3px;
  background-color: #e74c3c;
  color: white;
  cursor: pointer;
  position: absolute;
  top: 10px;
  right: 10px;
}
button:hover {
  background-color: #c0392b;
}
</style>
