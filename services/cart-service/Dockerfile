FROM golang:1.25.1-alpine AS build_shopping_cart_service
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app
# Copy shared module
COPY services/shared/ ./services/shared/
COPY services/cart-service/ ./services/cart-service/
WORKDIR /app/services/cart-service
RUN go mod download
RUN swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
RUN CGO_ENABLED=0 GOOS=linux go build \
  -ldflags="-s -w" \
  -o /app/cart-service \
  ./cmd/api

FROM alpine:latest AS shopping_cart_service
WORKDIR /app
COPY --from=build_shopping_cart_service /app/cart-service .
COPY --from=build_shopping_cart_service /app/services/cart-service/docs ./docs
RUN addgroup -g 1000 appuser && \
  adduser -D -u 1000 -G appuser appuser
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 8081
CMD ["./cart-service"]