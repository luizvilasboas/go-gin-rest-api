# go-gin-rest-api

This project was developed as part of the Go programming language training at Alura. It involves the creation of a REST API using `Gin` and `GORM`. Also, I'm using `testify` to test and mock the API. 

# Dependencies

To ensure smooth operation of serenatto, make sure you have the following dependencies installed:

- Go (version 1.21)
- Docker Compose (version 1.29)

# Usage

1. After installing `go`, clone the project and navigate to the project directory with the following commands:
```
git clone https://gitlab.com/alura-courses-code/golang/go-gin-rest-api.git
cd go-gin-rest-api
```

2. Install all dependencies required to run the project, especially `air`:
```
go mod tidy
```

3. Run the `docker-compose` to set up and host a `PostgreSQL` server in a container:
```
docker-compose up
```

4. Run the project and open `localhost:8080` in your browser using the following command:
```
air
```

## Contributing

If you wish to contribute to this project, feel free to open a merge request. We welcome all forms of contribution!

## License

This project is licensed under the [MIT](https://gitlab.com/alura-courses-code/golang/go-gin-rest-api/-/blob/main/LICENSE). Refer to the LICENSE file for more details.
