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

        <!-- Add Comment Form -->
        <div>
          <h4>Add a Comment:</h4>
          <input
              type="text"
              v-model="newCommentContent[blog.Id]"
              placeholder="Enter your comment"
          />
          <button @click="addComment(blog.Id)">Submit Comment</button>
        </div>
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
const newCommentContent = ref({}) // To store new comment content for each blog

const fetchBlogs = async () => {
  try {
    const response = await axios.post('http://localhost:8081/api/getAllBlogs', {
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
    const response = await axios.post('http://localhost:8081/api/getAllCommentsById', {
      blog_id: blog.Id
    })
    const data = response.data
    blog.comments = data.comments
  } catch (error) {
    console.error(`Error fetching comments for blog ID ${blog.Id}:`, error)
    blog.comments = []  // Ensure comments is initialized
  }
}

const addComment = async (blogId) => {
  const content = newCommentContent.value[blogId]
  if (!content) {
    alert('Comment content cannot be empty.')
    return
  }

  try {
    const response = await axios.post('http://localhost:8081/api/addCommentById', {
      blog_id: blogId,
      user_id: 1,  // Assuming user_id is 1; adjust as needed
      content: content
    })

    if (response.data.status === 'success') {
      alert('Comment added successfully!')

      // Add the new comment to the local comments array
      const blog = blogs.value.find(b => b.Id === blogId)
      if (blog) {
        blog.comments.push({
          Id: new Date().getTime(),  // Generate a temporary unique ID for the new comment
          Username: 'Current User',  // Replace with the actual username if available
          Content: content,
          CreatedAt: new Date().toISOString()
        })
      }

      // Clear the input field
      newCommentContent.value[blogId] = ''
    }
  } catch (error) {
    console.error(`Error adding comment to blog ID ${blogId}:`, error)
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
input[type="text"] {
  margin-right: 10px;
  padding: 5px;
  border-radius: 3px;
  border: 1px solid #ddd;
}
button {
  padding: 5px 10px;
  border-radius: 3px;
  border: none;
  background-color: #007bff;
  color: white;
  cursor: pointer;
}
button:hover {
  background-color: #0056b3;
}
</style>
