# [Gram] The sword is only as powerful as the person wielding it.

Inspired by the legendary Norse sword Gram, this Golang project wields the power
of simplicity and precision. A versatile tool for crafting robust and efficient
applications with mythical coding prowess.

## Prerequisites

-   [Go](https://golang.org/doc/install) ^1.21.6
-   [Docker](https://docs.docker.com/install/) ^18.09.2
-   [Docker Compose](https://docs.docker.com/compose/install/) ^1.23.2
-   [Serverless](https://www.serverless.com/framework/docs/getting-started) ^3.25.1
-   [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv1.html) ^1.16.14
-   [jq](https://stedolan.github.io/jq/download/) ^1.6.0
-   [air](https://github.com/cosmtrek/air) ^1.44.0
-   [go-swagger](https://goswagger.io/install.html#homebrewlinuxbrew) ^0.21.0

## Getting started

1. Initialize the app for the first time:
    ```
    make provision
    ```
2. Generate swagger API docs:
    ```
    make specs
    ```
3. Run the development server:
    ```
    make start
    ```

The application runs as a HTTP server at port 8083. You can log in as superadmin to the application by:

Then grab the `access_token` for authorization HTTP header:

```http
GET /v1/... HTTP/1.1
Host: localhost:8081
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ...
```

## Deployment

```
make deploy
```
