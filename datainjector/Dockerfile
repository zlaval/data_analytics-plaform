FROM golang:alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o datainjector ./cmd/datainjector
RUN chmod +x /app/datainjector

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/datainjector /app
CMD ["/app/datainjector"]