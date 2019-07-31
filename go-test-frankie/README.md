## Description

The test project is to build a simple RESTful service, based off a Swagger/OpenAPI v2.0 definition. JSON and YAML versions of the file are provided here: Swagger file

You will write the service using Golang

The service itself, is simple enough.

There is a single endpoint that takes a JSON payload.
The data to be passed in has simple validation rules that are described in the Swagger definition.
Successful validation returns a simple JSON structure
An error returns a description of all issues found.

## Build

````
docker build -t go-test-frankie .
````

## Run

````
docker run -p 8080:3200 go-test-frankie
docker run -d -p 8080:3200 go-test-frankie (as daemon)
