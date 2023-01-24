FROM golang:1.19-alpine AS builder
WORKDIR /
COPY ./services.boilerplate ./boilerplate
COPY ./services.shared ../services.shared
WORKDIR /boilerplate
RUN go mod download && go mod tidy
RUN go build -o main ./src/main.go

FROM alpine:latest AS runner

ENV PORT 8080

WORKDIR /app

COPY --from=builder /boilerplate/main .
COPY --from=builder /boilerplate/src/locales ./src/locales


EXPOSE $PORT

CMD ["/app/main"]
