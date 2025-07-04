# Stage 1: Build
FROM golang:1.22 as builder

WORKDIR /app
COPY . .

# Build the binary statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o userapi .

# Stage 2: Run
FROM scratch

COPY --from=builder /app/userapi /userapi

ENTRYPOINT ["/userapi"]
