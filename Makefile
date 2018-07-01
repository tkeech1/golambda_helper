test-local:
	docker build -t golambdahelper:latest . 
	docker run -it --rm -v ${PWD}:/go/src/github.com/tkeech1/golambdahelper -w /go/src/github.com/tkeech1/golambdahelper golambdahelper:latest go test