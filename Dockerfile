FROM golang:1.7

ENV REPO_PATH "$GOPATH/src/github.com/jonstacks/ilo"
RUN mkdir -p "$REPO_PATH"

COPY . $REPO_PATH

RUN go get "gopkg.in/xmlpath.v2"
RUN go install -v "github.com/jonstacks/ilo/cmd/ilo-sweep" && \
    go install -v "github.com/jonstacks/ilo/cmd/ilo-server"

CMD ["ilo-sweep"]
