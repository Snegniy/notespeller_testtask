#build stage
FROM golang:1.20-alpine AS builder

WORKDIR /notespeller
COPY . .
RUN go build -o app ./cmd/main.go

#run stage
FROM alpine
WORKDIR /notespeller
COPY --from=builder /notespeller/app .
COPY /public/swagger.json ./public/
COPY /.env .

EXPOSE 8000

ENTRYPOINT ["./app"]