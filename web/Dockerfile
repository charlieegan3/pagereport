FROM golang:latest
WORKDIR /app
RUN go get github.com/tools/godep
CMD godep go build -o web && /app/web 8081
