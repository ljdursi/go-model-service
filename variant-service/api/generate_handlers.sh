#!/bin/bash

# TODO make smarter by only re-generating files if they've been modified (makefile?)

# Generate handler utilities for the following resources: Individual, Variant
genny -in ./generics/generic_resource_utilities.go -out ./restapi/handlers/resource_utilities.go -pkg handlers gen "Resource=Individual,Variant"

# Generate POST handlers for the following resources: Individual, Variant
genny -in ./generics/generic_post.go -out ./restapi/handlers/post.go -pkg handlers gen "Resource=Individual,Variant"

# Generate GET (many) handlers for the following resources: Individual, Variant
genny -in ./generics/generic_get_many.go -out ./restapi/handlers/get_many.go -pkg handlers gen "Resource=Individual,Variant"