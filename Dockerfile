FROM golang:1.12.4-alpine3.9 as builder

RUN apk add --no-cache libc6-compat git gcc g++

WORKDIR /golang-demo

ADD . .

RUN go build -o golang-demo

FROM alpine:3.9

RUN apk add --no-cache libc6-compat

WORKDIR /

COPY --from=builder /golang-demo /

RUN chmod +x /golang-demo

CMD ["/golang-demo"]
