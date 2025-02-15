FROM golang:1.23 AS builder

ENV APP_PORT=8080

ENV CGO_ENABLED=0

WORKDIR /app
ADD . /app/


RUN go mod download && \
    go mod tidy && \
    go build -o v1/bin cmd/main.go


FROM alpine:3.21.3

WORKDIR /app

COPY --from=builder app/v1/ ./v1

COPY --from=builder app/resources ./resources

EXPOSE $APP_PORT

CMD ["./v1/bin"]