# Enube-challenge

A file importer created with Golang following best practices in software development, it also has some features such as user authentication generated tokens, pipelines for continuous integration and much more.

## Techs

* [Golang](https://go.dev/) - Programing language
* [Gin](https://github.com/gin-gonic/gin) - Gin Web Framework
* [Docker](https://www.docker.com/) - Containers
* [Docker hub](https://www.docker.com/) - Deploy Image application
* [Pipelines](https://docs.github.com/pt/actions) - CI/CD
* [Postgres](https://www.postgresql.org/) - Database
* [Jwt](https://www.postgresql.org/) - Token management for application access
* [Swaggo](https://github.com/swaggo) - Api docs
* [Raiwail](https://github.com/swaggo) - Deploy database
* [Zap](https://github.com/uber-go/zap) - Logs of the application

## **Endpoints**

### Auth
- **POST** `/api/v1/authentication/sign-in`: Log in a user and generate a token access

### Users

- **POST** `/api/v1/users/` : Create a new user
- **GET** `/api/v1/users/:email` : Get a user By email 

### Suppliers

- **POST** `/api/v1/suppliers/` : Import all suppliers from a file
- **GET** `/api/v1/suppliers` : Get all suppliers with pagination
- **GET** `/api/v1/suppliers/:id` : Get a suppliers by id

## Docs api

- **GET** `/api/v1/swagger/index.html` : See the docs


## Running this project

### Environment Setting

1. **Clone the repository**

 ```bash
 git clone https://github.com/seu-usuario/seu-repositorio.git
 cd enube-challenge
## Licence
This project is under license (MIT LICENCE) - see the file [LICENSE](https://github.com/hebertsanto/enube-challenge/blob/main/LICENCE) to more details
```

2 **Set the environment variables**

```
USER_DATABASE="your_user_database"
USER_PASSWORD="your_password"
DATABASE="your_database"
PORT="port_database"
HOST="host_databse"
SECRET_JWT="secret_jwt"
```

3 **Install dependencies**

```bash
go mod tidy
```

4 **Run the application**

```bash
go run cmd/api/main.go
```

## Pipeline

This project includes a pipeline configured to deploy the application to Docker Hub. The pipeline is configured to build and push the Docker image whenever there are changes to the repository.

```dockerfile
# Use the official Golang image as a parent image
FROM golang:1.22.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o enube-challenge ./cmd/api/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Define the command to run the application
CMD ["./enube-challenge"]

```

The entire deployment process is automated with a GitHub Action, streamlining integration and deployment to various services.