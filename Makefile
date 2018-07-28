test:
	go test -race -covermode=atomic
	
test-docker:
	docker build -t go_build:latest . 
	docker run -it --rm -v ${PWD}:/go/src/github.com/tkeech1/golambdahelper -w /go/src/github.com/tkeech1/golambdahelper go_build:latest /bin/bash -c "go get -v -t -d ./... && go test"

enable-go-modules:
	go mod -init -module github.com/tkeech1/golambdahelper

build:
	go build

# add modules or remove unnecessry ones
sync: 
	go mod -sync	

# if using 1.10 or earlier
get-deps:
	go get -v -t -d ./...
