# Stage 1: Build
FROM golang:1.22 as builder

WORKDIR /app
COPY . .

# ✅ Disable VCS stamping to fix the build error
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o userapi .

# Stage 2: Run
FROM scratch

COPY --from=builder /app/userapi /userapi

ENTRYPOINT ["/userapi"]
