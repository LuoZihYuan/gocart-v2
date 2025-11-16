FROM golang:1.25.1-alpine AS build_product_service
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app
# Copy shared module
COPY services/shared/ ./services/shared/
COPY services/product-service/ ./services/product-service/
WORKDIR /app/services/product-service
RUN go mod download
RUN swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
RUN CGO_ENABLED=0 GOOS=linux go build \
  -ldflags="-s -w" \
  -o /app/product-service \
  ./cmd/api

FROM alpine:latest AS product_service
WORKDIR /app
COPY --from=build_product_service /app/product-service .
COPY --from=build_product_service /app/services/product-service/docs ./docs
RUN addgroup -g 1000 appuser && \
  adduser -D -u 1000 -G appuser appuser
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 8080
CMD ["./product-service"]