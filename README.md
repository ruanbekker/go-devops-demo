# go-devops-demo

This project is a Go RESTful API application with a PostgreSQL (or Sqlite) backend and Prometheus Metrics, built with a structured and organized approach. 

## Project Structure

The project is organized into the following directories:

```bash
.
├── app
│   ├── cmd
│   │   └── server
│   │       └── main.go
│   ├── config
│   │   └── config.go
│   ├── controllers
│   │   └── user_controller.go
│   ├── models
│   │   ├── user_dto.go
│   │   └── user.go
│   ├── repositories
│   │   └── user_repository.go
│   ├── routers
│   │   └── router.go
│   ├── services
│   │   └── user_service.go
│   └── tests
│       └── user_test.go
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
└── README.md

11 directories, 15 files
```

## Folder Structure

- `app/`: Contains the main application code.
  - `cmd/`: Contains the entry point of the application.
    - `server/`: Contains the main entry point of the application.
      - `main.go`: The main file where the application starts.
  - `config/`: Contains configuration-related code.
    - `config.go`: Handles the database connection and configuration settings.
  - `controllers/`: Contains the HTTP handler functions for the API.
    - `user_controller.go`: Handles user-related HTTP requests.
  - `models/`: Contains the data models.
    - `user.go`: Defines the User model.
    - `user_dto.go`: Defines the User data object (only exposes specific fields).
  - `repositories/`: Contains the code for data access.
    - `user_repository.go`: Handles CRUD operations for the User model.
  - `routers/`: Contains the router setup.
    - `router.go`: Defines the API routes and connects them to the controllers.
  - `services/`: Contains the business logic.
    - `user_service.go`: Implements business logic for user operations.
  - `tests/`: Contains the test files.
    - `user_test.go`: Implements tests for user-related operations.

## Endpoints

- **GET /users**: Retrieves a list of all users.
- **GET /users/:id**: Retrieves a user by their ID.
- **POST /users**: Creates a new user.
- **PUT /users/:id**: Updates an existing user by their ID.
- **DELETE /users/:id**: Deletes a user by their ID.

## Resources

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [SQLite](https://www.sqlite.org/)

