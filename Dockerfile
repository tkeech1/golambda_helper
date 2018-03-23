FROM golang:1.9.2-stretch

RUN go get github.com/aws/aws-lambda-go/lambda
RUN go get github.com/aws/aws-lambda-go/events
RUN go get github.com/stretchr/testify/assert
RUN go get github.com/aws/aws-sdk-go
RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/aws/session
RUN go get github.com/aws/aws-sdk-go/service/dynamodb
RUN go get github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute