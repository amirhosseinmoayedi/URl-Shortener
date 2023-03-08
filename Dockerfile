FROM golang:1.19-bullseye as builder

WORKDIR /app

COPY go.* ./.

RUN CGO_ENABLED=0 go build -o url-shortner ./.

RUN chmod +x /app/url-shortner

FROM alpine:3.17.0

WORKDIR /app

COPY --from=builder /app /app

CMD ["url-shortner"]