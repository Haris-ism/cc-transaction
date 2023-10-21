FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
RUN apk add alpine-sdk
COPY . .
# RUN go mod download
RUN go build -o cc-transaction

FROM alpine:3.16

WORKDIR /app

RUN apk update && apk add --no-cache git
RUN apk update && apk add --no-cache tzdata
ENV TZ="Asia/Jakarta"

COPY --from=builder /app/cc-transaction .
COPY .env .

EXPOSE 9090

ENTRYPOINT [ "/app/cc-transaction" ]
