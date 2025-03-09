FROM golang:1.24.1 as builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

COPY ./migrations ./migrations

CMD ["./app"]