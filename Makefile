test:
	go get -v -t -d ./...
	go test -race -covermode=atomic
	
test-docker:
	docker build -t go_build:latest . 
	docker run -it --rm -v ${PWD}:/go/src/github.com/tkeech1/golambdahelper -w /go/src/github.com/tkeech1/golambdahelper go_build:latest /bin/bash -c "go get -v -t -d ./... && go test"