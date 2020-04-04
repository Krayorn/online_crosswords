FROM golang:latest

WORKDIR /go/src/online_crosswords

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "get", "golang.org/x/net/websocket"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ." -command="./online_crosswords"
