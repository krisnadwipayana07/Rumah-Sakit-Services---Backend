FROM golang:1.17.1-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o /app/main
EXPOSE 8080
CMD ["/app/main"]

FROM alpine:3.14
WORKDIR /root/
COPY --from=builder /app/config.json .
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
