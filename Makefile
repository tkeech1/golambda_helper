build:
	docker build -t golambda_helper:latest .
	docker run -it --rm -v ${PWD}:/go/src/golambda_helper -w /go/src/golambda_helper golambda_helper:latest \
		/bin/bash -c 'env GOOS=linux go build -ldflags="-s -w" -o bin/response_helper response_helper.go'	

test:
	docker build -t golambda_helper:latest . 
	docker run -it --rm -v ${PWD}:/go/src/golambda_helper -w /go/src/golambda_helper golambda_helper:latest go test

clean:
	docker rmi golambda_helper:latest
