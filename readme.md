# Go Todo App with JWT Authentication

This is a RESTful API for a Todo application built using Golang, the Gin framework, MongoDB for data storage, and JWT for user authentication.

## Features

- User registration and login
- JWT-based authentication
- CRUD operations for Todos
- MongoDB for data storage
- Middleware for JWT authentication

## Installation

### Prerequisites

- [Go](https://golang.org/) 1.18+
- [MongoDB](https://www.mongodb.com/) installed and running on `mongodb://localhost:27017`

### Setup

1. **Clone the Repository**

    ```bash
    git clone https://github.com/your-username/go-todo-app.git
    cd go-todo-app
    ```

2. **Install Dependencies**

    ```bash
    go get -u github.com/gin-gonic/gin
    go get go.mongodb.org/mongo-driver/mongo
    go get github.com/dgrijalva/jwt-go
    go get golang.org/x/crypto/bcrypt
    go get github.com/joho/godotenv
    ```

3. **Create a `.env` File**

    ```plaintext
    JWT_SECRET=your_secret_key
    ```

4. **Run the Application**

    ```bash
    go run main.go
    ```

## API Endpoints

### Authentication

- **Register**: `POST /register`
  - Request Body: `{ "username": "your_username", "password": "your_password" }`
  - Response: `{ "user_id": "<MongoDB User ID>" }`

- **Login**: `POST /login`
  - Request Body: `{ "username": "your_username", "password": "your_password" }`
  - Response: `{ "token": "<JWT Token>" }`

### Todos (Authenticated Routes)

These routes require the JWT token in the `Authorization` header as `Bearer <token>`.

- **Get All Todos**: `GET /todos`
- **Create a Todo**: `POST /todos`
  - Request Body: `{ "title": "your_title", "completed": false }`
- **Update a Todo**: `PUT /todos/:id`
  - Request Body: `{ "title": "new_title", "completed": true }`
- **Delete a Todo**: `DELETE /todos/:id`

## License

This project is licensed under the MIT License.