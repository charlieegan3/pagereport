FROM golang:latest
WORKDIR /usr/local/go/src/github.com/charlieegan3/pagereport/web
RUN ln -sf /usr/local/go/src/github.com/charlieegan3/pagereport/web /app
RUN go get github.com/tools/godep
CMD godep go build -o web && ./web 8081
