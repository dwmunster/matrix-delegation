FROM golang:1.16 as builder
WORKDIR /build
COPY . .
RUN go mod tidy \
    && CGO_ENABLED=0 GOOS=linux go build -trimpath -a -o matrix-delegation .

FROM alpine:3.13.5
WORKDIR /app
COPY --from=builder /build/matrix-delegation .
CMD ["/app/matrix-delegation"]
