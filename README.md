# Go Boilerplate Authentication Gateway

This project is an authentication gateway built in Go using the Gin framework. It provides user registration, login, and authenticated routes using JWT (JSON Web Token). The application uses PostgreSQL as the database and includes hot reload capabilities using Air.

## Features Implemented

- User Registration (Signup)
- User Login (JWT Token Generation)
- Protected Routes with JWT Authentication Middleware
- Password Hashing and Validation using bcrypt
- Environment Configuration using `.env` files
- Hot Reloading during development with Air

## Requirements

- Go 1.17+
- PostgreSQL
- `Air` for hot reloading (for development purposes)

## Setup Instructions

### Step 1: Clone the Repository

```sh
git clone https://github.com/yourusername/go-boilerplate.git
cd go-boilerplate
```

### Step 2: Install Dependencies

Install the required Go packages:

```sh
go mod tidy
```

### Step 3: Set Up Environment Variables

Create a `.env` file in the root directory of your project with the following content:

```env
DB_DSN=host=localhost user=psql_admin password=admin dbname=auth_gateway_db port=5432 sslmode=disable
PORT=3000
JWT_SECRET=your_secret_key
GIN_MODE=debug
```

- **DB_DSN**: PostgreSQL connection string
- **PORT**: The port the server will listen on
- **JWT_SECRET**: Secret key for signing JWT tokens
- **GIN_MODE**: Set to `release` for production, `debug` for development

### Step 4: Database Setup

Ensure PostgreSQL is installed and running. Create a database for the project:

```sql
CREATE DATABASE auth_gateway_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

### Step 5: Install Air for Hot Reloading (Optional)

Install Air to enable hot reloading during development:

1. **Download Air**:

   ```sh
   curl -fLo ~/.local/bin/air https://raw.githubusercontent.com/air-verse/air/master/bin/air
   chmod +x ~/.local/bin/air
   ```

2. **Add Air to PATH** (if not already done):

   ```sh
   export PATH=$PATH:~/.local/bin
   ```

3. **Initialize Air Configuration**:

   ```sh
    air init
   ```

### Step 6: Run the Application

To start the server:

```sh
air
```

Alternatively, if not using `Air`, you can run:

```sh
go run main.go
```

The server will start on the port specified in the `.env` file (default: `3000`).

## API Documentation

### 1. User Registration (Signup)

- **Endpoint**: `POST /signup`
- **Description**: Registers a new user with their first name, last name, email, and password.
- **Request Body**:
  ```json
  {
      "firstName": "John",
      "lastName": "Doe",
      "email": "john.doe@example.com",
      "hashedPassword": "password123"
  }
  ```
- **Response**:
  ```json
  {
      "message": "User created successfully"
  }
  ```

### 2. User Login

- **Endpoint**: `POST /login`
- **Description**: Logs in the user and provides a JWT token for authentication.
- **Request Body**:
  ```json
  {
      "email": "john.doe@example.com",
      "hashedPassword": "password123"
  }
  ```
- **Response**:
  ```json
  {
      "token": "<JWT_TOKEN>"
  }
  ```

### 3. Access Protected Route (Welcome)

- **Endpoint**: `GET /welcome`
- **Description**: Access a protected route. Requires a valid JWT token in the Authorization header.
- **Authorization Header**:
  ```
  Authorization: Bearer <JWT_TOKEN>
  ```
- **Response**:
  ```json
  {
      "message": "Welcome john.doe@example.com"
  }
  ```

## Folder Structure

```
/go-boilerplate
  ├── main.go
  ├── /controllers
  │     └── user_controller.go
  ├── /middlewares
  │     └── auth_middleware.go
  ├── /models
  │     └── user.go
  ├── /routes
  │     └── routes.go
  ├── /services
  │     └── user_service.go
  ├── /utils
  │     └── password_utils.go
  ├── .env
  ├── .air.toml
  └── go.mod
```

## Development Tools and Libraries Used

- **Gin**: Web framework used to handle HTTP requests.
- **GORM**: ORM for interacting with PostgreSQL.
- **JWT (dgrijalva/jwt-go)**: Used for creating and validating JSON Web Tokens.
- **bcrypt**: For password hashing and validation.
- **Air**: Used for hot reloading during development.
- **Godotenv**: Load environment variables from `.env` files.

## How to Test

### Testing with cURL

1. **Signup**:
   ```sh
   curl --location --request POST 'http://localhost:3000/signup' \
   --header 'Content-Type: application/json' \
   --data-raw '{
       "firstName": "John",
       "lastName": "Doe",
       "email": "john.doe@example.com",
       "hashedPassword": "password123"
   }'
   ```

2. **Login**:
   ```sh
   curl --location --request POST 'http://localhost:3000/login' \
   --header 'Content-Type: application/json' \
   --data-raw '{
       "email": "john.doe@example.com",
       "hashedPassword": "password123"
   }'
   ```

3. **Access Protected Route**:
   ```sh
   curl --location --request GET 'http://localhost:3000/welcome' \
   --header 'Authorization: Bearer <JWT_TOKEN>'
   ```

Replace `<JWT_TOKEN>` with the token received from the login response.

## Known Issues and Troubleshooting

- **JWT Token Invalid**: Ensure that the secret key (`JWT_SECRET`) is consistent across the application. The key used during login must be the same as the key used for validation.
- **Nil Pointer Dereference**: If you see a nil pointer dereference, ensure that the `db` instance is properly set in the service layer using `SetDB()`.
- **Hot Reload Not Working**: If `Air` is not reloading automatically, ensure that `.air.toml` is correctly configured and that `air` is in your PATH.

## Future Enhancements

- Add refresh token functionality for extended login sessions.
- Implement role-based access control (RBAC) for different user roles.
- Add more detailed error handling and validation for user inputs.
- Implement rate limiting for authentication routes.