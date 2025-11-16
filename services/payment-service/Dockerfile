FROM golang:1.25.1-alpine AS build_payment_service
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app
# Copy shared module
COPY services/shared/ ./services/shared/
COPY services/payment-service/ ./services/payment-service/
WORKDIR /app/services/payment-service
RUN go mod download
RUN swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
RUN CGO_ENABLED=0 GOOS=linux go build \
  -ldflags="-s -w" \
  -o /app/payment-service \
  ./cmd/api

FROM alpine:latest AS payment_service
WORKDIR /app
COPY --from=build_payment_service /app/payment-service .
COPY --from=build_payment_service /app/services/payment-service/docs ./docs
RUN addgroup -g 1000 appuser && \
  adduser -D -u 1000 -G appuser appuser
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 8083
CMD ["./payment-service"]