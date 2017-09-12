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

## Testing

In order to run all tests, simply type `go test`. Some useful flags you should
consider:
* `-v`: verbose mode, e.g. lists test methods
* `-cover`: gives you test coverage statistics
* `-run`: Specify test cases  and subtests that you want to run
  * `-run ''`: Run all tests
  * `-run Root`: Run top-level tests matching "Root", e.g. "TestRootHandler"
  * `-run Sine/Input=float`: Run top-level tests matching "Sine", run subtests
  matching "Input=float"

Pitfalls:
* Test methods have to be public.
* A parent test will only complete once all of its subtests complete.

## Managing dependencies

Dependency management has been a hot topic in the Go community for a while now.
Rephrasing the discussion here would be useless, instead refer to this
[post]{https://blog.gopheracademy.com/advent-2016/saga-go-dependency-management/}.


Right now, there are multiple packages that serve the same purpose:
* [godep]{https://github.com/tools/godep}: This is the tool that will be merged
into the standard toolchain. The project is inactive since the repo was moved
(see below). The tool is working and ready for production use.
* [dep]{https://github.com/golang/dep}: This tool can be used now and will be
merged into the toolchain with Go 1.10.
* Other tools that are used in the community are [gb]{https://getgb.io/},
[Glide]{https://github.com/Masterminds/glide}, etc.


`godep` is the most production-ready tool right now. The workflow is as follows:
1. Install a dependency via `go get`, e.g. `go get github.com/pkg/errors`.
2. Save dependencies to a project-based `vendor` directory using `godep save`.
3. Packages can be found in the `vendor` directory and all dependencies are
listed in the `Godeps/Godeps.json` file.
4. In order to update a package, run `go get -u foo/bar` and then
`godep update foo/bar`.

*TODO: include information on the dep tool!*
