FROM golang:latest

RUN mkdir -p /app

RUN go get "github.com/gin-gonic/gin"

RUN go get "gopkg.in/mgo.v2"

RUN go get -v "github.com/spf13/viper"

WORKDIR /app

ADD . /app

RUN go build ./server.go

CMD [ "./server"]