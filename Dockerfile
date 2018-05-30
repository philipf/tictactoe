FROM golang:alpine

WORKDIR /go/

RUN apk update && apk add git
RUN go get -d -v github.com/philipf/tictactoe
RUN go install -v ./...

CMD ["tictactoe"]
