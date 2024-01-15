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


RUN make build-producer

#final stage
FROM alpine:3.18.5

COPY --from=build-env /app/bin/producer /usr/local/bin/producer

EXPOSE 8080
CMD ["producer"]
