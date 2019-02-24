FROM golang
RUN mkdir -p /go/src/github.com/k-yomo/tweet_scheduler
WORKDIR /go/src/github.com/k-yomo/tweet_scheduler
COPY . .
RUN go get github.com/labstack/echo &&\
    go get github.com/lib/pq &&\
    go get github.com/jinzhu/gorm &&\
    go get github.com/dgrijalva/jwt-go &&\
    go get gopkg.in/yaml.v2 &&\
    go get github.com/joho/godotenv &&\
    go get github.com/jinzhu/configor
CMD ["go", "run", "/go/src/github.com/k-yomo/tweet_scheduler/server.go"]
