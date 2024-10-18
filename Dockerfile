FROM golang:1.23

WORKDIR /go/project/apod

ADD go.mod go.sum main.go Makefile ./
ADD internal ./internal
ADD docs ./docs
ADD www ./www

EXPOSE 8080

CMD ["go", "run", "main.go"]