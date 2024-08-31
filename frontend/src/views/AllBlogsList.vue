<template>
  <div>
    <h2>All Blogs {{ authorId }}</h2>
    <ul v-if="blogs.length > 0">
      <li v-for="blog in blogs" :key="blog.Id">
        <h3>{{ blog.Username }} (ID: {{ blog.Id }})</h3>
        <p>{{ blog.Content }}</p>
        <small>Created At: {{ formatDate(blog.CreatedAt) }}</small>

        <!-- Render Comments -->
        <div v-if="blog.comments && blog.comments.length > 0">
          <h4>Comments:</h4>
          <ul>
            <li v-for="comment in blog.comments" :key="comment.Id">
              <p><strong>{{ comment.Username }}:</strong> {{ comment.Content }}</p>
              <small>Created At: {{ formatDate(comment.CreatedAt) }}</small>
            </li>
          </ul>
        </div>
        <p v-else>Loading comments or no comments available.</p>
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
    const response = await axios.post('http://localhost:8081/getAllBlogs', {
      userId: 1
    })
    const data = response.data
    authorId.value = data.authorId
    blogs.value = data.blogs

    // Fetch comments for each blog
    for (const blog of blogs.value) {
      await fetchCommentsForBlog(blog)
    }
  } catch (error) {
    console.error('Error fetching blogs:', error)
  }
}

const fetchCommentsForBlog = async (blog) => {
  try {
    const response = await axios.post('http://localhost:8081/getAllCommentsById', {
      blog_id: blog.Id
    })
    const data = response.data
    blog.comments = data.comments
  } catch (error) {
    console.error(`Error fetching comments for blog ID ${blog.Id}:`, error)
    blog.comments = []  // Ensure comments is initialized
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
h4 {
  margin-top: 10px;
}
</style>
