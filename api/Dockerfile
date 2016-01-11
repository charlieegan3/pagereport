FROM ianneub/go:latest
ADD . /usr/local/go/src/github.com/charlieegan3/pagereport/api
WORKDIR /usr/local/go/src/github.com/charlieegan3/pagereport/api
RUN go get sourcegraph.com/sourcegraph/go-selenium
RUN go build
CMD ./api 8080
