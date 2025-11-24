# ---- Build stage ----
FROM golang:1.24-alpine AS builder
WORKDIR /avito/backend

# Копируем зависимости
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Копируем весь код
COPY backend .

# Собираем бинарник
RUN go build -o main ./cmd

# ---- Final stage ----
FROM alpine:3.18
WORKDIR /avito
COPY --from=builder /avito/backend/main ./main
RUN apk add --no-cache ca-certificates

EXPOSE 8080
CMD ["./main"]
