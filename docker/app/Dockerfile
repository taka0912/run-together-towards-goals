#Step 1 Builld
FROM golang:alpine as builder

WORKDIR /go/src/app/
ADD src/ .

RUN apk update \
    && apk add --no-cache git vim less curl make gcc \
    && go mod download \
    && go build -o app main.go

#Step2 Move build file
FROM alpine:3.9

RUN apk add --no-cache mariadb-client mysql-dev alpine-sdk

COPY --from=builder /go/src/app/ /app

ENV PORT=${PORT}

WORKDIR /app

CMD ["./app"]
#CMD ash -c "ash start_app.sh && /bin/ash"
