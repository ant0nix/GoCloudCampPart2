FROM golang:latest

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-connection.sh

RUN go mod download
RUN go build -o goCLoudTask ./cmd/main.go

CMD ["./goCLoudTask"]