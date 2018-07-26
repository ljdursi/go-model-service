# Go Model Service

Based on Jonathan Dursi's [OpenAPI variant service demo](https://github.com/CanDIG/openapi_calls_example), this toy service demonstrates the go-swagger/pop stack with CanDIG API best practices.

[![Build Status](https://travis-ci.org/CanDIG/go-model-service.svg?branch=master)](https://travis-ci.org/CanDIG/go-model-service)

## Stack

- [Sqlite3](https://www.sqlite.org/index.html) database backend
- [Go](https://golang.org/) (Golang) backend
- Gobuffalo [pop](https://github.com/gobuffalo/pop) is used as a pseudo-orm for interfacing between the go code and the sqlite3 database. The `soda` CLI auto-generates boilerplate Go code for models and migrations, as well as performing migrations on the database. `fizz` files are used for defining database migrations in a Go-like syntax (see [syntax documentation](https://github.com/markbates/pop/tree/master/fizz).)
- [Go-swagger](https://goswagger.io/) auto-generates boilerplate Go code from a `swagger.yml` api definition. [Swagger](https://swagger.io/) tooling isbased on the [OpenAPI](https://www.openapis.org/) specification.
- [dep](https://golang.github.io/dep/docs/introduction.html) is used for dependency management.
- Gobuffalo [validate](https://github.com/gobuffalo/validate) is a framework used for writing custom validators. Some of their validators in the `validators` package are used as-is.

## Installation

1. [Install Go](https://golang.org/doc/install). Make sure to set up the `$PATH` and `$GOPATH` environment variables according to [these instructions](https://www.digitalocean.com/community/tutorials/), and to understand the expected contents of the three `$GOPATH` subdirectories: `$GOPATH/src`, `$GOPATH/pkg`, and `$GOPATH/bin`.
2. [Install dep](https://golang.github.io/dep/docs/installation.html)
3. [Install go-swagger](https://goswagger.io/install.html) (releases 0.15.0 or later strongly recommended.)
4. [Install pop](https://github.com/gobuffalo/pop). See the [Unnoficial Pop Book](https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/installation.html) for instructions. Make sure to include sqlite3 support with `tags sqlite` in your installation commands, as follows:
  ```
  $ go get -u -v -tags sqlite github.com/gobuffalo/pop/...
  $ go install -tags sqlite github.com/gobuffalo/pop/soda
  ```
5. Install this service as desired:
  * Into a local directory of your choosing with
  `$ git checkout https://github.com/CanDIG/go-model-service.git`
  * Into your go environment with 
  `$ go get github.com/CanDIG/go-model-service`
  (installation directory will be `$GOPATH/src`).
6. In the root directory of this project (ie. the directory where `Gopkg.lock` and `Gopkg.toml` are found), run 
`$ dep ensure`
to install all project import dependencies in the new `vendor` directory.

## Running The Service

Navigate to the root directory of the project and enter the following commands to start the server on port 3000 (or any other port; modify the parameter in the last line as desired):
```
$ cd variant-service/api
$ go build -tags sqlite ./variant-service/api/cmd/project-server/main.go
$ ./main --port=3000
```
