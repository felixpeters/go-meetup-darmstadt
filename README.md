# Go Meetup Darmstadt: Test, Build, Run

This repo contains the code accompanying my talk called "Go Basics: Test, Build,
Run", delivered on September 14, 2017 in Darmstadt, Germany.

## The app

Our app contains two basic views:
1. A landing page with usage instructions at the root URL.
2. A page which displays the result of the sine operation for a given input.

Inputs for the sine operation can be delivered via the query parameter *input*:
* **Example query:** `http://localhost:8080/sine?input=5`

If the input cannot be parsed to a floating-point number, an error will be
shown.

## Building

TODO: Build the app locally and in the Go workspace. Name executable.

## Testing

TODO: Test handler using httptest package.

TODO: Test mathematical function using table-driven tests, subtests.

TODO: Write another test using a testing package.

## Running

TODO: Write Makefile for running tests, building and running app inside a Docker
container.

## Managing dependencies

TODO: Initialize `/vendor` directory using `godep` and save dependencies.
