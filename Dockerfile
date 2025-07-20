FROM golang:1.24.5-alpine3.22 as builder
RUN apk add --no-cache build-base
WORKDIR /src
COPY . .
ENV CGO_ENABLED=1
RUN go build -o bot cmd/main.go

FROM alpine:3.22
COPY --from=builder /src/bot /usr/bin/bot
EXPOSE 8080
ENTRYPOINT ["/usr/bin/bot"]
