FROM golang:1.10.5-alpine3.8

WORKDIR /go/src/app
COPY maingo.go .
RUN go build -o main .
EXPOSE 8000
CMD [“./main”]

// docker build -t appgo:1.0 .
// docker images
// docker run -d –name appgo -p 8000:8000 appgo:1.0
// docker ps