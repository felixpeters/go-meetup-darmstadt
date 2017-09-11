test:
	go test

compile:
	go build -o bin/app

install:
	go install

build-image:
	@echo "### Compiling"
	GOOS=linux go build -o bin/app
	@echo "### Building image"
	docker build -t go-meetup-darmstadt .

run-docker: build-image
	docker run -it -p 8080:8080 go-meetup-darmstadt
