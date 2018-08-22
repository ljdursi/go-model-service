#!/bin/bash

# This script installs the service by auto-generating the boilerplate code necessary
# for running the service.
# It also creates a new development database to write into.

export POP_PATH=$GOPATH/src/github.com/CanDIG/go-model-service/model-vs/config

# Use the dep tool to install all project import dependencies in a
# new vendor directory.
# It is run vendor-only to avoid modification of Gopkg.lock, which
# contains information about vital sub-packages for go-swagger that
# can not be explicitly constrained in Gopkg.toml.
# For example, package "github.com/go-openapi/runtime/flagext" is
# required by go-swagger but is *not* solved into Gopkg.lock if
# `dep ensure` is run prior to `swagger validate`. Therefore it is
# important to read the existing Gopkg.lock file in the initial
# installation, rather than solve for a new one.
# For more information, see: https://golang.github.io/dep/docs/ensure-mechanics.html
dep ensure -vendor-only

# Create a sqlite3 development database and migrate it to the schema
# defined in the model-vs/data directory, using the soda tool from pop
cd model-vs/config
soda create -c ./database.yml -e development
soda migrate up -c ./database.yml -e development -p ../data/migrations
cd ../..

# Swagger generate the boilerplate code necessary for handling API requests
# from the model-vs/api/swagger.yml template file
cd model-vs/api
swagger generate server -A variant-service swagger.yml # This will generate a server named variant-service, which is important for maintaining compatibility with the configure_variant_service.go middleware configuration file.
cd ../..

# Run a script to generate resource-specific request handlers for middleware,
# from the generic handlers defined in the model-vs/api/generics package,
# using the CanDIG-maintained CLI tool genny
cd model-vs/api
./generate_handlers.sh
pwd
ls
ls model-vs/api/restapi/handlers/get_many.go
cat model-vs/api/restapi/handlers/get_many.go
cd ../..

# Now that all the necessary boilerplate code has been auto-generated, compile the server
go build -tags sqlite -o ./main model-vs/api/cmd/variant-service-server/main.go
