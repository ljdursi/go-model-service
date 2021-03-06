#!/bin/bash

# TODO make smarter by only re-generating files if they've been modified (makefile?)

apiDir=$GOPATH/src/github.com/CanDIG/go-model-service/model-vs/api

# Generate handler utilities for the following resources: Individual, Variant, Call
genny -in $apiDir/generics/generic_resource_utilities.go -out $apiDir/restapi/utilities/resource_utilities.go -pkg utilities gen "Individual=Individual,Variant,Call"

# Generate POST handlers for the following resources: Individual, Variant, Call
genny -in $apiDir/generics/generic_post.go -out $apiDir/restapi/handlers/post.go -pkg handlers gen "Individual=Individual,Variant,Call"

# Generate GET (one) handler for the following resources: Individual, Variant, Call
genny -in $apiDir/generics/generic_get_one.go -out $apiDir/restapi/handlers/get_one.go -pkg handlers gen "Individual=Individual,Variant,Call"

# Generate GET (many) handlers for the following resources: Individual, Variant, Call
genny -in $apiDir/generics/generic_get_many.go -out $apiDir/restapi/handlers/get_many.go -pkg handlers gen "Individual=Individual,Variant,Call"

# Generate GET (by another resource) handlers for the following resources: {Individuals by Variant with junction in Call}, {Variants by Individual with junction in Call}
# There are presently issues with the genny tool that prevent this step from working properly.
# Please see github.com/CanDIG/genny/issues to help resolve some of these issues.
# genny -in $apiDir/generics/generic_get_by.go -pkg handlers gen "Individual=Individual Variant=Variant Call=Call" > $apiDir/restapi/handlers/get_by.go
# genny -in $apiDir/generics/generic_get_by.go -pkg gen "Individual=Variant Variant=Individual Call=Call" >> $apiDir/restapi/handlers/get_by.go

# Generate transformers for the following resources: Individual, Variant, Call
genny -in $apiDir/generics/generic_transformations.go -out $apiDir/restapi/handlers/transformations.go -pkg handlers gen "Individual=Individual,Variant,Call"