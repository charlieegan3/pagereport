FROM ianneub/go:latest
ADD . /usr/local/go/src/github.com/charlieegan3/pagereport/
WORKDIR /usr/local/go/src/github.com/charlieegan3/pagereport/
RUN go get sourcegraph.com/sourcegraph/go-selenium
RUN go build
CMD ./pagereport 8080
