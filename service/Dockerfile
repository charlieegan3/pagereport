FROM iron/go:dev

RUN go get sourcegraph.com/sourcegraph/go-selenium

ENV SRC_DIR=/go/src/github.com/charlie/pagereport/service
WORKDIR $SRC_DIR

ADD . $SRC_DIR

RUN go build -o service

CMD ./service $PORT
