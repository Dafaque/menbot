FROM golang:1.23.4-alpine3.21 as builder
RUN apk add --no-cache build-base
WORKDIR /src
COPY . .
ENV CGO_ENABLED=1
RUN go build -o bot cmd/main.go

FROM alpine:3.21
COPY --from=builder /src/bot /usr/bin/bot
ENTRYPOINT ["/usr/bin/bot"]
