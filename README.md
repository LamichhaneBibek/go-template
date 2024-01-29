# Golang Backend Boilerplate
## Overview
 
This Golang backend boilerplate provides a foundation for building web applications using the Gin framework, PostgreSQL as the database, and Zero Logger for logging. It aims to streamline the initial setup process, allowing developers to focus on building features rather than configuring the project structure and dependencies.

## Features

-	[Gin Framework](Gin Framework): A fast and lightweight web framework for Go.
-   [PostgreSQL Database](PostgreSQL Database): Utilizes PostgreSQL as the database for data storage.
-   [Zero Logger](Zero Logger): A logging library for Go, designed for simplicity and ease of use.

## Getting Started
	Prerequisites

	- Go
	- PostgreSQL

## Installation
1. Clone the repo:
```sh
git clone https://github.com/LamichhaneBibek/go-template.git
```

2. Change directory:
```sh
cd go-template
```

3. Install dependencies:
```sh
go mod tidy
```

4. Set up the PostgreSQL database. Create a new database and update the database configuration in the config/config.go file.

5. Run the application:
```sh
make server
```
or 
```sh
go run cmd/main.go
```

The application should be accessible at http://localhost:8080.

## Configuration

    Database Configuration: Update the PostgreSQL database connection details in the config/config.go file.

## Project Structure

```
./
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── api.go
│   │   ├── handlers
│   │   │   └── user_handlers.go
│   │   ├── helpers
│   │   │   └── response_helpers.go
│   │   ├── middleware
│   │   ├── models
│   │   │   ├── base_models.go
│   │   │   └── user_models.go
│   │   ├── routes
│   │   │   └── user_routes.go
│   │   └── validations
│   ├── common
│   ├── config
│   │   ├── config.go
│   │   └── config.yml
│   ├── constants
│   │   └── constants.go
│   ├── db
│   │   ├── database.go
│   │   └── migrate.go
│   ├── services
│   │   ├── token_services.go
│   │   └── user_services.go
│   ├── utils
│   │   ├── error.go
│   │   └── logger
│   │       ├── category.go
│   │       ├── logger.go
│   │       └── zero_logger.go
│   └── views
│       └── user_views.go
├── Makefile
└── README.md
```

## Contributing

Feel free to open issues and pull requests. Contributions are welcome!

## License

Distributed under the MIT License. See `LICENSE` for more information.

Happy Coding!