FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-backend ./cmd/main.go

FROM scratch
COPY --from=builder /app .
EXPOSE 8080
EXPOSE 5432
CMD ["./go-backend"]