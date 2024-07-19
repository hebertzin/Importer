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

## **Endpoints**

- **POST** `/api/v1/authentication/sign-in`: Log in a user and generate a token access
- **POST** `/api/v1/users/` : Create a new user
- **GET** `/api/v1/users/:email` : Get a user By email 

## Licence

This project is under license (MIT LICENCE) - see the file [LICENSE](https://github.com/hebertsanto/enube-challenge/blob/main/LICENCE) to more details
