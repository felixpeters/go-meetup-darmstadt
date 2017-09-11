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

We can build our app in several different ways:
* Use `go build` in order to build an executable with the repository name in 
the repository directory.
* Use `go build -o bin/app` in order to build an executable with the name *app* in
the `bin` directory.
* Use `go install` in order to build an executable which can be found in your
`$GOPATH/bin` directory.

## Testing

TODO: Test handler using httptest package.

TODO: Test mathematical function using table-driven tests, subtests.

TODO: Write another test using a testing package.

## Running

In order to test your app, you can simply run your main package using the
command `go run main.go`.


If you built the app locally (see above), you can run it using the command 
`bin/app`.


If you built the app using `go install` and your `$GOPATH/bin` directory is in
your `PATH`, you can also run your app using the command `go-meetup-darmstadt`.


With Go it is possible to compile code on your platform (e.g. macOS) for usage
on another platform (e.g. Linux). This is known as *cross-compiling* and 
becomes very useful if you want to run your app e.g. in a Docker container.
Compiling on your system and just putting the executable into a small base
image allows for Docker images with extremely small file sizes. For more
information check out the Dockerfile.

## Managing dependencies

TODO: Initialize `/vendor` directory using `godep` and save dependencies.
