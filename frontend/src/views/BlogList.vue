<template>
  <div>
    <h2>Blogs by Author {{ authorId }}</h2>
    <ul v-if="blogs.length > 0">
      <li v-for="blog in blogs" :key="blog.Id">
        <h3>{{ blog.Username }} (ID: {{ blog.Id }})</h3>
        <p v-if="!isEditing(blog.Id)">{{ blog.Content }}</p>
        <div v-else>
          <textarea v-model="editingContent" rows="3" style="width: 100%;"></textarea>
          <button @click="updateBlog(blog.Id)">Save</button>
          <button @click="cancelEditing">Cancel</button>
        </div>
        <small>Created At: {{ formatDate(blog.CreatedAt) }}</small>
        <button @click="deleteBlog(blog.Id)">Delete</button>
        <button @click="editBlog(blog.Id, blog.Content)" v-if="!isEditing(blog.Id)">Edit</button>
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
const editingBlogId = ref(null)
const editingContent = ref('')

const fetchBlogs = async () => {
  try {
    const response = await axios.post('http://localhost:8081/api/GetAllBlogsById', {
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
    const response = await axios.post('http://localhost:8081/api/deleteBlog', {
      blogId: blogId
    })
    if (response.data.status === 'success') {
      blogs.value = blogs.value.filter(blog => blog.Id !== blogId)
    } else {
      console.error('Failed to delete blog.')
    }
  } catch (error) {
    console.error('Error deleting blog:', error)
  }
}

const editBlog = (blogId, content) => {
  editingBlogId.value = blogId
  editingContent.value = content
}

const cancelEditing = () => {
  editingBlogId.value = null
  editingContent.value = ''
}

const isEditing = (blogId) => {
  return editingBlogId.value === blogId
}

const updateBlog = async (blogId) => {
  try {
    const response = await axios.post('http://localhost:8081/api/modifyBlogById', {
      BlogId: blogId,
      Content: editingContent.value
    })
    if (response.data.status === 'success') {
      const blog = blogs.value.find(blog => blog.Id === blogId)
      if (blog) blog.Content = editingContent.value
      cancelEditing()
    } else {
      console.error('Failed to update blog.')
    }
  } catch (error) {
    console.error('Error updating blog:', error)
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
  margin-right: 5px;
}
button:hover {
  background-color: #c0392b;
}
textarea {
  margin-bottom: 10px;
}
</style>
