Here's a README format for user authentication, based on the description you provided:

---

# Authentication API with JWT and PostgreSQL

This project is a simple **User Authentication API** built using **Go**, **JWT** (JSON Web Tokens) for authentication, and **PostgreSQL** for storing user data.

## Prerequisites

Before you can run this API and interact with it via `curl`, you'll need to:

- Set up **PostgreSQL** on your local machine
- Run the **Go** server
- Use **curl** to test the various endpoints

### Prerequisites:

- **Go (Golang)** installed on your machine
- **PostgreSQL** installed and running
- **curl** for testing the API endpoints

## Setup

### 1. Set Up PostgreSQL

#### Install PostgreSQL

If you don’t have PostgreSQL installed, follow the installation instructions for your operating system:

- [PostgreSQL installation guide](https://www.postgresql.org/download/)

#### Create a Database and User

1. Access PostgreSQL by running the following command:

   ```bash
   psql postgres
   ```

   This will log you into the PostgreSQL prompt.

2. Create a new database:

   ```sql
   CREATE DATABASE mydb;
   ```

3. Create a new user (replace `yourusername` and `yourpassword` with your desired username and password):

   ```sql
   CREATE USER yourusername WITH PASSWORD 'yourpassword';
   ```

4. Grant privileges to the user on the `mydb` database:

   ```sql
   GRANT ALL PRIVILEGES ON DATABASE mydb TO yourusername;
   ```

5. Exit PostgreSQL:

   ```sql
   \q
   ```

#### Modify the Connection String in the Go Application

In the `main.go` file of your Go project, ensure that the connection string matches your PostgreSQL configuration:

```go
// Connection string format
connString := "postgres://yourusername:yourpassword@localhost:5432/mydb?sslmode=disable"
```

### 2. Run the Go Application

1. Install Go dependencies (in the project directory):

   ```bash
   go mod tidy
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

   This will start the server on `http://localhost:8080`.

## API Endpoints

The API supports three main endpoints:

1. **POST `/signup`** – Register a new user
2. **POST `/login`** – Log in and get a JWT token
3. **GET `/protected`** – Access a protected route, requires a valid JWT token

## Using curl to Test the API

### 1. Signup Endpoint

To register a new user, send a `POST` request to `/signup` with `username` and `password` as JSON data.

#### curl Command:

```bash
curl -X POST http://localhost:8080/signup \
     -H "Content-Type: application/json" \
     -d '{"username": "newuser", "password": "password123"}'
```

This will register a new user with the username `newuser` and password `password123`.

#### Response:

```json
{
    "message": "User registered successfully!"
}
```

---

### 2. Login Endpoint

To log in and get a JWT token, send a `POST` request to `/login` with the `username` and `password`.

#### curl Command:

```bash
curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"username": "newuser", "password": "password123"}'
```

If the login is successful, you will receive a JWT token.

#### Response:

```json
{
    "token": "your.jwt.token"
}
```

---

### 3. Access Protected Route

To access a protected route, include the JWT token obtained from the `/login` endpoint in the `Authorization` header.

#### curl Command:

```bash
curl -X GET http://localhost:8080/protected \
     -H "Authorization: Bearer your.jwt.token"
```

Replace `your.jwt.token` with the JWT token you received from the `/login` response.

#### Response:

```json
{
    "message": "Hello, newuser! You have accessed a protected route."
}
```

---

