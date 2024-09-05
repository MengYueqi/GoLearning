<template>
  <div>
    <h2>Add a New Blog</h2>
    <form @submit.prevent="submitBlog">
      <div>
        <label for="authorId">Author ID:</label>
        <input type="number" v-model="authorId" required />
      </div>
      <div>
        <label for="content">Content:</label>
        <textarea v-model="content" required></textarea>
      </div>
      <button type="submit">Submit Blog</button>
    </form>
    <p v-if="statusMessage">{{ statusMessage }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

const authorId = ref('')
const content = ref('')
const statusMessage = ref('')

const submitBlog = async () => {
  try {
    const response = await axios.post('http://localhost:8081/api/addBlog', {
      AuthorId: parseInt(authorId.value),
      Content: content.value
    })
    if (response.data.status === 'success') {
      statusMessage.value = 'Blog submitted successfully!'
      // 清空表单
      authorId.value = ''
      content.value = ''
    } else {
      statusMessage.value = 'Failed to submit the blog.'
    }
  } catch (error) {
    console.error('Error submitting blog:', error)
    statusMessage.value = 'An error occurred while submitting the blog.'
  }
}
</script>

<style scoped>
h2 {
  margin-bottom: 20px;
}
form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
label {
  margin-bottom: 5px;
}
input,
textarea {
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
  width: 300px;
}
button {
  width: 100px;
  padding: 10px;
  border: none;
  border-radius: 5px;
  background-color: #007bff;
  color: white;
  cursor: pointer;
}
button:hover {
  background-color: #0056b3;
}
p {
  margin-top: 20px;
  color: green;
}
</style>
