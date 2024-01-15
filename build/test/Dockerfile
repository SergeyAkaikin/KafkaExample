#build stage
FROM golang:1.21-alpine AS build-env


RUN apk --no-cache update && \
    apk add --no-cache git make gcc libc-dev

ENV CGO_ENABLED 1
ENV GOOS=linux
ENV GOARCH=amd64


WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .

RUN go test -tags musl -c -o ./bin/test ./test/produce_consume_test.go

#final stage
FROM alpine:3.18.5

COPY --from=build-env /app/bin/test /usr/local/bin/test

EXPOSE 8080
CMD ["test"]