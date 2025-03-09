FROM golang:1.24.1 as builder

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

FROM alpine:latest

COPY --from=builder /app .

COPY ./migrations /migration

CMD ["./app"]