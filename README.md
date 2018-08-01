# Go Model Service

Based on Jonathan Dursi's [OpenAPI variant service demo](https://github.com/CanDIG/openapi_calls_example), this toy service demonstrates the go-swagger/pop stack with CanDIG API best practices.

[![Build Status](https://travis-ci.org/CanDIG/go-model-service.svg?branch=master)](https://travis-ci.org/CanDIG/go-model-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/candig/go-model-service)](https://goreportcard.com/report/github.com/candig/go-model-service)

## Stack

- [Sqlite3](https://www.sqlite.org/index.html) database backend
- [Go](https://golang.org/) (Golang) backend
- Gobuffalo [pop](https://github.com/gobuffalo/pop) is used as an orm-like for interfacing between the go code and the sqlite3 database. The `soda` CLI auto-generates boilerplate Go code for models and migrations, as well as performing migrations on the database. `fizz` files are used for defining database migrations in a Go-like syntax (see [syntax documentation](https://github.com/markbates/pop/tree/master/fizz).)
- [Go-swagger](https://goswagger.io/) auto-generates boilerplate Go code from a `swagger.yml` API definition. [Swagger](https://swagger.io/) tooling isbased on the [OpenAPI](https://www.openapis.org/) specification.
- [dep](https://golang.github.io/dep/docs/introduction.html) is used for dependency management.
- Gobuffalo [validate](https://github.com/gobuffalo/validate) is a framework used for writing custom validators. Some of their validators in the `validators` package are used as-is.

## Installation

Prior to installing new programs, run `$ which <program-name>` to check if it is already installed on your machine. If there is a blank output rather than a path to the program binary, it needs to be installed.

1. [Install Go](https://golang.org/doc/install). Make sure to set up the `$PATH` and `$GOPATH` environment variables according to [these instructions](https://www.digitalocean.com/community/tutorials/), and to understand the expected contents of the three `$GOPATH` subdirectories: `$GOPATH/src`, `$GOPATH/pkg`, and `$GOPATH/bin`.
2. [Install gcc](https://gcc.gnu.org/install/).
3. [Install sqlite3](https://www.tutorialspoint.com/sqlite/sqlite_installation.htm).
4. [Install dep](https://golang.github.io/dep/docs/installation.html)
5. Install this go-model-service as desired:
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

## For Developers

### Installing Dev Tools

In addition to all steps in the Installation section above, install the following tools:
1. [Install go-swagger](https://goswagger.io/install.html) (release 0.15.0 strongly recommended.)
2. [Install pop](https://github.com/gobuffalo/pop). See the [Unnoficial Pop Book](https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/installation.html) for instructions. Make sure to include sqlite3 support with `tags sqlite` in your installation commands, as follows:
  ```
  $ go get -u -v -tags sqlite github.com/gobuffalo/pop/...
  $ go install -tags sqlite github.com/gobuffalo/pop/soda
  ```

### Instructions for Development

#### Go-Swagger

Go-Swagger is a tool that automatically generates the boilderplate Go code needed to build a server, based on the API definitions for the service being provided by that server.

See [goswagger.io](https://goswagger.io/) for installation instructions, tutorials, use-cases, etc. If you find yourself having trouble with the installation, check the [prerequisites](https://goswagger.io/generate/requirements.html). The [Todo List Tutorial](https://goswagger.io/tutorial/todo-list.html) (Simple Server) is a good place to start if you've never used the go-swagger before. 

Go-Swagger uses Swagger 2.0, which is based on the OpenAPI specification. "Swagger" and "OpenAPI" are often used interchangeably, which can be confusing when trying to learn the pertinent tool or set up your environment. See [this post](https://swagger.io/blog/api-strategy/difference-between-swagger-and-openapi/) for an explanation of the relationship between the two.

##### Generating The Server

The API definitions are written by the developer in a `swagger.yml` file. To validate that the `swagger.yml` file follows the specification, run
`$ swagger validate <path-to-target-swagger.yml>`.

To auto-generate a server based on the entities and endpoints described in the `swagger.yml` file, run
`$ swagger generate server -A <server-name> <path-to-target-swagger.yml>.`
For example, for this service, from the `go-model-service/variant-service/api` directory, you would run the following to re-generate the server:
`$ swagger generate server -A variant-service swagger.yml`

The backend can now be implemented by modifying the endpoint handlers in `restapi/configure_<server-name>.go`. The connection to the data backend is made in these handlers. Other configuration such as middleware setup is also written in this file, in its respective methods.

##### Boilerplate Code and Directory Structure

All files in the api directory are auto-generated (and auto-replaced upon calling
`$ swagger generate server <path-to-target-swagger.yml>`) except for the following:
- swagger.yml: The swagger definition.
- configure_variant_service.go: Auto-generated but safe to edit.
- main: Generated by calling `$ go build cmd/variant-service-server/main.go`

The auto-generated boilerplate code includes:
- models
The API-facing models for entities are generated into the `models` package.
Models are generated from the `definitions` defined in the `swagger.yml` file.
- endpoints
The endpoints for the API are defined in `paths` in the `swagger.yml` file, and from their definition the `operations` package is populated with endpoint parameters, validation, responses, URL building, etc.
However the backend handlers for these endpoints, ie. what is done with the received request, must be written manually in confifure_variant_service.go, in the configureAPI method. By default, the handlers return 501: Not Implemented responses.
- server
The go server files and main.go are auto-generated.
- configuration
The `configure_variant_service.go` file is auto-generated but safe to edit. This is where the manually written backend goes, where requests are handled following their automatic transformation into go stucts, and where responses/payloads are assigned.
Middleware can be plugged in here.
The connection to the data backend/memory store (ie. the ORM and/or database) should be made here.

#### GoBuffalo Pop

Pop is an ORM-like that is used to interface between a go backend and one of several database languages. 

See the [pop README](https://github.com/gobuffalo/pop#pop--) for installation and use instructions. There is also an [Unofficial pop Book](https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/) with tutorials, [Quick Start](https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/installation.html) being a good place to begin.

Since the database used in this project is `sqlite3`, there are slight modifications that must be made to some commands in the form of a `-tags sqlite` option. These are detailed in the [Installing CLI Support](https://github.com/gobuffalo/pop#installing-cli-support) section of the Pop documentation.

##### Pop Migrations

Soda is a CLI tool for generating pop migration files and models, as well as for running up- and down-migrations. Migrations are described in `.fizz` or `.sql` files, and beyond simple migrations such as adding/dropping columns, these files must be manually populated with explicit migration instructions.

Fizz provides a Go-like syntax for writing migrations, but [you may instead opt for writing SQL migrations as desired](https://github.com/gobuffalo/pop#generating-migrations). The fizz syntax is described [here]](https://github.com/markbates/pop/tree/master/fizz).

###### Migrating Pop Models

Soda can generate models in pop from command-line input, but these models must be manually edited when migrations cause modifications to the database table that a model corresponds to.

For example, if you add a `province` column to the `individual` table in a migration, the `individual.go` model must have that field added to its `type Individual struct`. You may also want to add validations for this new field in the `Validate` method of the `individual.go` model.

##### Validating Pop Models

`Validate` method contained in each model Go file is called upon each `ValidateAndSave` (or similar) call, to ensure that the data about to be entered in the database meets developer-defined constraints.

The `validators` package from Gobuffalo [validate](https://github.com/gobuffalo/validate) is the only set of validators automatically imported by Pop, but this `validate` framework allows for the creation of custom validators as needed. See the `tools/validators` directory for a simple example, or  the Unofficial pop Book's tutorial on [Writin Validations](https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/advanced-topics/writing-validations.html) for a more complex example.

##### Handling Nulls With Pop

There is some complexity introduced in representing database tables with Go structs. Since Go only allows `nil` values for pointers, one must employ a work-around for handling nulls retrieved from the database, which is particularly necessary for validating their presence or absence.

By default, null-values from the database are transformed into Go's zero-values for a column that is represented in Go by a non-nillable type. For example, the `Chromosome` field of the `Variant` model (in `variant.go`) is of type string, and if a value for `chromosome` is not supplied in an entry, the value of the `Chromosome` field is `""`. The `validators` package used to validate *required* fields in pop only checks for these fields having a non-zero value.

This project uses the `pop/nulls` package to handle non-nullable fields that should be permitted to have zero values, such as the `Start` field of the `Variant` model. This field is of type nulls.Int, which is able to support null values (and therefore a custom `int_is_not_null.go` validator.)