# go-gin-rest-api

> A REST API built with Go, Gin, and GORM, developed as part of an Alura training course.

## About the Project

This project is a REST API developed to practice Go, using the Gin framework for routing and the GORM library for database interaction. It also includes unit tests and mocks using the Testify library. The entire environment is containerized with Docker Compose.

## Tech Stack

*   [Go](https://golang.org/)
*   [Gin](https://gin-gonic.com/)
*   [GORM](https://gorm.io/)
*   [Testify](https://github.com/stretchr/testify) (for testing)
*   [PostgreSQL](https://www.postgresql.org/)
*   [Docker Compose](https://docs.docker.com/compose/)
*   [Air](https://github.com/cosmtrek/air) (for live reloading)

## Usage

Below are the instructions for you to set up and run the project locally.

### Prerequisites

You need to have the following software installed:

*   [Go](https://golang.org/dl/) (version 1.21 or higher)
*   [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
*   [Air](https://github.com/cosmtrek/air#installation) (for development)

### Installation and Setup

Follow the steps below:

1.  **Clone the repository**
    ```bash
    git clone https://github.com/luizvilasboas/go-gin-rest-api.git
    ```

2.  **Navigate to the project directory**
    ```bash
    cd go-gin-rest-api
    ```

3.  **Install Go dependencies**
    ```bash
    go mod tidy
    ```

4.  **Start the database container**
    ```bash
    docker-compose up -d
    ```

### Workflow

To run the project with live reloading, use `air`:
```bash
air
```
The API will be available at `http://localhost:8080`.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
