FROM golang:1.16-alpine as builder
RUN apk update && apk upgrade && \
    apk --update add git make bash

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
RUN make build

FROM alpine

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata

WORKDIR /app
EXPOSE 8080
COPY --from=builder /app/bin/engine /app/engine

CMD /app/engine
