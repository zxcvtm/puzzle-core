FROM golang:1.14

WORKDIR /go/src/app
COPY . .
RUN go mod init
RUN go get  -v ./...
RUN go install -v ./...
RUN go build -o main .
CMD ./main