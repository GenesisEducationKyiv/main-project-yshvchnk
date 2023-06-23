FROM golang:1.20-alpine as build-stage

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /bitcoin-service .

FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

# Copy the binary and the .env file
WORKDIR /app
COPY --from=build-stage /bitcoin-service /app/bitcoin-service
COPY --from=build-stage /app/.env /app/.env

EXPOSE 8080

CMD ["./bitcoin-service"]
