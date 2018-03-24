test-local:
	docker build -t golambda_helper:latest . 
	docker run -it --rm -v ${PWD}:/go/src/github.com/tkeech1/golambda_helper -w /go/src/github.com/tkeech1/golambda_helper golambda_helper:latest go test