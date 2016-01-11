FROM golang:latest
ADD . /usr/local/go/src/github.com/charlieegan3/pagereport/web
WORKDIR /usr/local/go/src/github.com/charlieegan3/pagereport/web
RUN go build
CMD ./web 8081
