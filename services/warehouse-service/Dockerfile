FROM golang:1.25.1-alpine AS build_warehouse_service
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app
# Copy shared module
COPY services/shared/ ./services/shared/
COPY services/warehouse-service/ ./services/warehouse-service/
WORKDIR /app/services/warehouse-service
RUN go mod download
RUN swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
RUN CGO_ENABLED=0 GOOS=linux go build \
  -ldflags="-s -w" \
  -o /app/warehouse-service \
  ./cmd/api

FROM alpine:latest AS warehouse_service
WORKDIR /app
COPY --from=build_warehouse_service /app/warehouse-service .
COPY --from=build_warehouse_service /app/services/warehouse-service/docs ./docs
RUN addgroup -g 1000 appuser && \
  adduser -D -u 1000 -G appuser appuser
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 8082
CMD ["./warehouse-service"]