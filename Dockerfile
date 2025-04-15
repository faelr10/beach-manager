# Etapa 1 - build
FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

# Etapa 2 - imagem final
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 10000

CMD ["./main"]
