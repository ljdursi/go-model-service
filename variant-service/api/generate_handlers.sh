#!/bin/bash

# TODO make smarter by only re-generating files if they've been modified (makefile?)
# TODO rename file?

# Generate handler utilities for the following resources: Individual, Variant, Call
genny -in ./generics/generic_resource_utilities.go -out ./restapi/handlers/resource_utilities.go -pkg handlers gen "Resource=Individual,Variant,Call"

# Generate POST handlers for the following resources: Individual, Variant, Call
genny -in ./generics/generic_post.go -out ./restapi/handlers/post.go -pkg handlers gen "Resource=Individual,Variant,Call"

# Generate GET (one) handler for the following resources: Individual, Variant, Call
genny -in ./generics/generic_get_one.go -out ./restapi/handlers/get_one.go -pkg handlers gen "Resource=Individual,Variant,Call"

# Generate GET (many) handlers for the following resources: Individual, Variant, Call
genny -in ./generics/generic_get_many.go -out ./restapi/handlers/get_many.go -pkg handlers gen "Resource=Individual,Variant,Call"

# Generate transformers for the following resources: Individual, Variant, Call
genny -in ./generics/generic_transformers.go -out ../transformations/transformers.go -pkg transformations gen "Resource=Individual,Variant,Call"