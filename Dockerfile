FROM golang
RUN mkdir -p $GOPATH/src/github.com/k-yomo/tweet_scheduler
WORKDIR $GOPATH/src/github.com/k-yomo/tweet_scheduler
COPY . .
RUN go get github.com/labstack/echo &&\
    go get github.com/lib/pq &&\
    go get github.com/jinzhu/gorm &&\
    go get github.com/dgrijalva/jwt-go &&\
    go get gopkg.in/yaml.v2
CMD ["$GOPATH/src/github.com/k-yomo/tweet_scheduler/server.go"]
