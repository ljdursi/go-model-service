#!/bin/bash

# This script prepares the environment for the install.sh script by installing
# the necessary CLI tools. It does not install libraries imported by the
# program's code itself; those are installed by the go dep tool into the
# vendor package.
# This file is used by .travis.yml and the Dockerfile, but is not recommended
# for local use as some of these programs may already be installed on
# your system.
# Please see the README for "Installing the Stack" for installation instructions
# for these tools.

# Install sqlite3 (database backend); sudo if necessary
apt-get update || sudo apt-get update
apt-get install -y sqlite3 libsqlite3-dev || sudo apt-get install -y sqlite3 libsqlite3-dev

# Install dep (managing project dependencies from Go import statements)
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Install Go-swagger (code-gen of boilerplate server Go code from OpenAPI definition)
curl -o $GOPATH/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/0.16.0/swagger_linux_amd64
chmod +x $GOPATH/bin/swagger

# Install pop (ORM-like for interfacing with the database backend)
# Note the use of `-tags sqlite` in the install statement for the soda CLI
go get -u -v -tags sqlite github.com/gobuffalo/pop/...
go install -tags sqlite github.com/gobuffalo/pop/soda # soda is the pop CLI

# Install genny (code-gen solution for generics in Go)
go get github.com/CanDIG/genny
which genny