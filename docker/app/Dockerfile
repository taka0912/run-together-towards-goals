FROM golang:latest

WORKDIR /go/src/app/
COPY src/ .


#RUN apt-get update -qq \
#    && apk add-get -y git vim mariadb-client \
##    && go mod download \
##    && go build -o /server \
#    && go get -u github.com/gin-gonic/gin \
#    && go get github.com/jinzhu/gorm \
#    && go get github.com/go-sql-driver/mysql \
#
##   / && export CGO_ENABLED=0

RUN go get -u github.com/gin-gonic/gin \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql \
    && apt-get update \
    && apt-get install -y mariadb-client

CMD ["sh", "start_app.sh"]
