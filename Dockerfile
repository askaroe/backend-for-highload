FROM golang:1.22.1 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o demo-app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN ls -la

COPY --from=builder /app/demo-app .
COPY --from=builder /app/db/migrations ./db/migrations

# Command to run the executable
CMD ["./demo-app"]
