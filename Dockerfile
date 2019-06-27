FROM golang:latest

WORKDIR /go/src/parser
COPY ./app/parser.go .

RUN go get github.com/urfave/cli

RUN go install -v .

CMD ["parser"]
