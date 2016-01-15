FROM golang:latest
ADD . /usr/local/go/src/github.com/charlieegan3/pagereport/web
WORKDIR /usr/local/go/src/github.com/charlieegan3/pagereport/web
RUN go get github.com/tools/godep
RUN godep go build
CMD ./web 8081
