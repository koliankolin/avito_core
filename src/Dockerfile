FROM golang:latest

WORKDIR /go/src/parser
COPY ./app/* ./

RUN go get github.com/urfave/cli
RUN go get github.com/essentialkaos/translit
RUN go get github.com/PuerkitoBio/goquery

RUN go install -v .


#ENTRYPOINT parser
#CMD ["get"]

