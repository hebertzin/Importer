# Enube-challenge

Enube-Challenge is a file import application developed in Golang. This project follow good software development practices and features robust functionalities, including user authentication with JWT tokens, CI/CD pipelines for seamless integration, and data management using PostgreSQL. Designed to handle and import supplier data, the application provides endpoints for user authentication, creation and retrieval of user information, and supplier management.

## Techs

* [Golang](https://go.dev/) - Programing language
* [Gin](https://github.com/gin-gonic/gin) - Gin Web Framework
* [Docker](https://www.docker.com/) - Containers
* [Docker hub](https://www.docker.com/) - Deploy Image application
* [Pipelines](https://docs.github.com/pt/actions) - CI/CD
* [Postgres](https://www.postgresql.org/) - Database
* [Jwt](https://www.postgresql.org/) - Token management for application access
* [Swaggo](https://github.com/swaggo) - Api docs
* [Railway](https://railway.app/) - Deploy database
* [Zap](https://github.com/uber-go/zap) - Logs of the application
* [Postman](https://www.postman.com/) - Testing and documenting APIs.

## **Endpoints**

### Authentication
- **POST** `/api/v1/authentication/sign-in`: Log in a user and generate a token access

### Users

- **POST** `/api/v1/users/` : Create a new user
- **GET** `/api/v1/users/:email` : Get a user By email 

### Suppliers

- **POST** `/api/v1/suppliers/` : Import all suppliers from a file
- **GET** `/api/v1/suppliers` : Get all suppliers with pagination
- **GET** `/api/v1/suppliers/:id` : Get a suppliers by id

## Docs api (Github pages)

[![Texto alternativo](https://github.com/user-attachments/assets/59a1f914-0c72-421a-8a98-91eece7c3c17)](https://hebertsanto.github.io/Enube-challenge/)



## Running this project

### Environment Setting

1. **Clone the repository**

 ```bash
 git clone https://github.com/hebertsanto/Enube-challenge
 cd Enube-challenge
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

## Design patterns

1 [Dependency injection](https://www.freecodecamp.org/news/a-quick-intro-to-dependency-injection-what-it-is-and-when-to-use-it-7578c84fa88f/#:~:text=In%20software%20engineering%2C%20dependency%20injection,be%20used%20(a%20service).)

2 [Dependency inversion](https://medium.com/@tbaragao/solid-d-i-p-dependency-inversion-principle-e87527f8d0be)

3 [Data Transfer Objects](https://docs.abp.io/en/abp/latest/Data-Transfer-Objects)

## Postman collection

- [Click here to postman collection](https://www.postman.com/descent-module-architect-9422719/workspace/enube/collection/36500427-d1489007-cff7-4dd2-8d6f-f869f3c7462d)

## Conclusion

I really enjoyed doing this challenge, I learned a lot of important concepts about the language and new ways of solving problems, it was a great experience.
