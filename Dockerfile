FROM golang:latest

COPY . /microplate

WORKDIR /microplate

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULE=auto go build -ldflags "-w -extldflags -static" -o main main.go

CMD ["./main"]