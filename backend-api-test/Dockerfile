FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . . 

RUN go build -o backend ./cmd/app

FROM golang:1.22-alpine AS runner

WORKDIR /app

COPY --from=builder /app/backend /app/backend

EXPOSE 8080

CMD ["/app/backend"]