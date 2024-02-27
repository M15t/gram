# GRAM - Golang API starter kit & project templates

GRAM is a Golang starter kit for developing API backend services. It is designed to help you kickstart your project, skipping the 'setting-up part' and jumping straight to writing business logic.

GRAM follows SOLID principles, with package design being inspired by several package designs, including Ben Johnson's [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1), [Go Standard Package Layout](https://github.com/golang-standards/project-layout) with my own ideas applied to both. The idea for building this project inspired by [GORSK v2](https://github.com/ribice/gorsk).

This starter kit currently provides:

- Full serverless stack with Lambda and Aurora
- Full featured API endpoints for authentication, view profile, changing password and CRUD operations on the user entity
- Application configuration via .env files and AWS Parameter Store with multi-environments supported
- JWT authentication
- Full featured RBAC powered by Casbin
- Common HTTP security headers, and CORS
- Request marshaling and data validation
- Structured error response handling
- API Docs using SwaggerUI
- Go Modules supported
- Easy to extend and/or delete, services are decoupled by using interface{}
- Great performance
- Mocking using stdlib (TODO)
- Complete test coverage, containerized database query tests (TODO)
- And more!

The application runs as an HTTP server at port 8080. It provides the following API endpoints:

- `GET /swagger-ui/` (with trailing slash): launches SwaggerUI in browser
- `POST /auth/login`: accepts email/passwords and returns jwt access token
- `POST /auth/refresh-token`: refreshes sessions and returns new jwt access token
- `GET /admin/users`: returns list of users
- `GET /admin/users/:id`: returns single user
- `POST /admin/users`: creates a new user
- `PATCH /admin/users/:id`: updates a user
- `DELETE /admin/users/:id`: deletes a user
- `GET /admin/users/me`: returns info about currently logged in user
- `PATCH /admin/users/me/password`: changes password for logged in user

## Project Structure

- Root directory contains things not related to code directly, e.g. docker-compose, CI/CD, readme, .env, go.mod, makefile, etc...
- `cmd` directory contains code for starting applications (main packages). The directory name for each application should match the name of the executable you want to have.
- `config` package contains the Configuration struct which represents the application settings, with Load func exported for loading the settings from various sources (AWS Parameter Store, ENV vars, file).
- `deploy` contains all the packages, tools, configs related to the deployment progress
- `docker` contains Docker files and stuff to build custom Docker image.
- `internal` directory contains rest of the code. The packages inside this directory are meant for internal use only. Some recommended subpackages:
  - `api` naming corresponds to `cmd/api` folder naming, contains multiple folders for each domain it interacts with, for example: user, car, appointment etc. Domain directories, like user, contain all application/business logic as well as http transport and other platforms if needed
  - `db` contains database connection implements to connect to designated database server
  - `model` contains domain types for the application. This package only contains simple data types and interfaces for reference
  - `repo` contains repository interfaces and implementations for data models
  - `tools` contains tooling packages
- `third_party` same ideas as `internal` here above but contains packages to be shareable with other external packages. Other projects will import these libraries expecting them to work, so think twice before you put something here!
- `scripts` contains scripts to perform various build, install, analysis, etc operations. These scripts keep the root level Makefile small and simple.

## Getting started

To use GRAM as a starting point of a real project whose package name is something like `awesome`:

1. Clone this repo to anywhere you want. Open it with VSCode!
2. Do a global replacement (CMD+SHIFT+H) with this search `gram`, replace to `awesome`.
3. Do another global replacement with this search `gram`, replace to `awesome` with "Preserve Case" (AB) option checked.
4. This template is set up with PostgreSQL by default. If you want to switch to MySQL database, search for `EnableMySQL` label and follow instructions.
5. Update the configurations in `config/config.go` and `.env.example` files according to your needs.
6. Update the `deploy/**/serverless.yml` files as needed.
7. Remove the `.git` directory, this file and other files that not being used as well. Init new git repo and push your first commit!
8. You're all set! Follow the steps in [README.md](README.md) file to initialize the project.

## Implementing CRUD of another table or object

Let's say you have a table named 'cars' that handles user's cars. To implement CRUD on this table you need:

1. Inside `internal/model`, create a new file named `car.go` then put your struct, and struct's methods if needed.
2. Create a new repository file named `car.go` in the `internal/repo/gormrepo` directory, refer to the other repository implementation.
   - Remember to run `go generate ./...` to update the UnitOfWork implementation once done.
3. Inside `internal/api` or specific `internal/api/{service}`, create a new folder named `car`, then create these files inside:
   - `http.go` contains HTTP handlers and API specs
   - `service.go` contains the init function and all interfaces that required by this package
   - `car.go` which named as the package's name, contains all the logic of the package
   - And order file to seperate the logic if needed
4. In `cmd/api/main.go`, wire up all the logic, by initializing car repository, services, and passing it to the http service afterwards.
5. Create a new migration item in `cmd/migration/main.go`, copy the new `Car` struct into it and use `tx.AutoMigrate()` to create the table in database.
6. Run `make migrate` to apply the change to database.
7. Run `make specs` to update API specs.
8. Run `make start` and test your new API at `http://localhost:8080/swagger-ui/`
