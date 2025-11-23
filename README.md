# Event Management REST API

A simple RESTful API for managing events and user registrations, built with Go (Gin) and PostgreSQL.

## Features

- **User Authentication**: Signup and Login with JWT authentication.
- **Event Management**: Create, Read, Update, and Delete (CRUD) events.
- **Event Registration**: Users can register and unregister for events.
- **Database**: Persists data using PostgreSQL.

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL
- **Driver**: [lib/pq](https://github.com/lib/pq)

## Prerequisites

- Go installed (version 1.18+ recommended)
- PostgreSQL installed and running

## Setup & Installation

1.  **Clone the repository**:

    ```bash
    git clone <repository-url>
    cd <project-directory>
    ```

2.  **Install Dependencies**:

    ```bash
    go mod tidy
    ```

3.  **Database Configuration**:

    - Ensure you have a PostgreSQL database created (default name expected is `test_project`).
    - Open `db/db.go` and update the connection string with your PostgreSQL credentials:
      ```go
      // db/db.go
      conn, err := sql.Open("postgres", "postgres://postgres:<your_password>@localhost:5432/test_project?sslmode=disable")
      ```
    - _Note: In a production environment, use environment variables for sensitive credentials._

4.  **Run the Application**:
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

## API Endpoints

### Authentication

| Method | Endpoint  | Description           | Payload                                                    |
| :----- | :-------- | :-------------------- | :--------------------------------------------------------- |
| `POST` | `/signup` | Register a new user   | `{"email": "user@example.com", "password": "password123"}` |
| `POST` | `/login`  | Login and receive JWT | `{"email": "user@example.com", "password": "password123"}` |

### Events

| Method   | Endpoint      | Description          | Auth Required | Payload                                                                                               |
| :------- | :------------ | :------------------- | :------------ | :---------------------------------------------------------------------------------------------------- |
| `GET`    | `/events`     | Get all events       | No            | -                                                                                                     |
| `GET`    | `/events/:id` | Get a specific event | No            | -                                                                                                     |
| `POST`   | `/events`     | Create a new event   | **Yes**       | `{"name": "Event Name", "description": "...", "location": "...", "dateTime": "2025-01-01T10:00:00Z"}` |
| `PUT`    | `/events/:id` | Update an event      | **Yes**       | `{"name": "Updated Name", ...}`                                                                       |
| `DELETE` | `/events/:id` | Delete an event      | **Yes**       | -                                                                                                     |

### Registrations

| Method   | Endpoint               | Description           | Auth Required |
| :------- | :--------------------- | :-------------------- | :------------ |
| `POST`   | `/events/:id/register` | Register for an event | **Yes**       |
| `DELETE` | `/events/:id/register` | Cancel registration   | **Yes**       |

## Project Structure

- `main.go`: Application entry point.
- `handlers/`: Contains API route handlers (controllers).
- `models/`: Defines data structures and database operations.
- `db/`: Database connection and table initialization.
- `middlewares/`: Authentication middleware.
- `utils/`: Utility functions (hashing, JWT).

## Testing

You can test the API using tools like Postman, Insomnia, or `curl`.
A sample python script `python.py` (if available) can also be used to verify the flow.
