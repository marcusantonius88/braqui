FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /build/server ./cmd/braqui/

FROM alpine:3.20

RUN apk --no-cache add ca-certificates wget

WORKDIR /app
COPY --from=builder /build/server .

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

CMD ["./server"]
