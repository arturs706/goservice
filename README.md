
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