# GoLearning
The record of Go learning.

## 1. LeetCode Hot 100

## 2. Gin

We develop a web backend by using the framework -- **Gin**. Gin is a high-performance, lightweight Go language web framework, known for its simple and easy-to-use routing and middleware design, fast request processing speed, and good scalability. On the backend, we use **gorm** to access and operate **MySQL**.
On the frontend, we used VUE3 to build a simple frontend page to display backend data.

The following operations were performed on the database using gorm:

<details><summary>Operations</summary>
<table>
    <thead>
        <tr>
            <th>API</th>
            <th>Database Operation</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Login</td>
            <td>Select</td>
        </tr>
        <tr>
            <td>GetAllBlogsById</td>
            <td>Select</td>
        </tr>
        <tr>
            <td>addBlog</td>
            <td>Insert</td>
        </tr>
        <tr>
            <td>deleteBlog</td>
            <td>Delete</td>
        </tr>
    </tbody>
</table>
</details>

