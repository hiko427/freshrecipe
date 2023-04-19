# Build stage
FROM golang:1.20.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
# FROM alpine:3.17
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY app.env .
# COPY templates .
# COPY public .

EXPOSE 8080
CMD [ "/app/main" ]