
# User Service
This repository contains a user service built using Go and the Fiber web framework. The service is structured following the principles of Clean Architecture to ensure maintainability, scalability, and testability.

Application Structure
The project is organized into several layers, each with a distinct responsibility:

```
src
├── domain
│   └── users.go
├── infrastructure
│   ├── db
│   │   └── sql.go
│   ├── auth
│   │   └── jwt.go
│   ├── security
│   │   └── password.go
│   └── router
│       └── router.go
├── interface
│   ├── controllers
│   │   └── user-controller.go
│   └── repository
│       ├── user-repository.go
│       └── repository.go
├── main.go
└── usecases
    └── user-usecase.go
```

## Layers Explanation

### domain
- This layer contains the core business logic and entities. It is independent of other layers, ensuring that the business rules are isolated.
  - `users.go`: Defines the user entity and business rules related to users.

### infrastructure
- This layer provides implementations for interfaces defined in the domain and usecases layers. It includes external services, database access, and framework integrations.
  - `db`: Database connection and SQL-related logic.
    - `sql.go`: Handles database interactions and queries.
  - `auth`: Authentication mechanisms.
    - `jwt.go`: JWT (JSON Web Token) generation and validation.
  - `security`: Security utilities.
    - `password.go`: Password hashing and validation.
  - `router`: HTTP routing setup.
    - `router.go`: Defines the API routes and connects them to the controllers.

### interface
- This layer contains the external interfaces and controllers for handling incoming requests and outgoing responses.
  - `controllers`: HTTP controllers for managing user-related requests.
    - `user-controller.go`: Handles HTTP requests for user operations.
  - `repository`: Interfaces for data storage and retrieval.
    - `user-repository.go`: Implements user-related database operations.
    - `repository.go`: General repository interfaces and implementations.

### usecases
- This layer contains application-specific business rules. It acts as a bridge between the domain and the interface layers.
  - `user-usecase.go`: Contains use cases for user operations such as creating, updating, and retrieving users.

### main.go
- The entry point of the application. It initializes the application, sets up the infrastructure, and starts the server.

## Getting Started

### Prerequisites
- Go
- Fiber

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/arturs706/user-service.git

### Install dependencies:
```bash
go mod tidy
```

### Set up environment variables:
- Create a `.env` file in the root directory and add necessary environment variables.

### Running the Application
- To start the application, run:
  ```bash
  go run main.go
    ```
- The application will start on `localhost:2001` by default.

## Project Goals
- This project aims to:
  - Demonstrate the use of Clean Architecture in a Go Fiber application.
  - Provide a scalable and maintainable structure for building web services.
  - Showcase best practices for Go development, including dependency injection, SOLID principles, and unit testing.
