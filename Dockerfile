FROM golang:latest

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate.linux-amd64 /usr/local/bin/migrate


RUN chmod +x wait-connection.sh

RUN go mod download
RUN go build -o goCLoudTask ./cmd/main.go

CMD ["./goCLoudTask"]