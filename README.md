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
### Technology Stack

#### Gin

Use Gin to complete the backend, process data according to the set API and return the results.

#### Vue3

Use Vue3 to complete the front-end design, request data from the back-end according to the route and receive data. Render the received data and provide users with a convenient interface to operate the data.

#### MySQL

Use MySQL to store data persistently and provide it to the backend interface for adding, deleting, modifying and querying data.

#### Redis

Use Redis to cache hot data. The backend first searches for data in Redis to reduce the pressure on the MySQL database.

#### JWT

Use JWT technology to authenticate users. The front end receives the JWT token and processes and saves it. The JWT token is then carried in the request and authenticated by the back end. Data is returned only if it meets the conditions to ensure data security.

