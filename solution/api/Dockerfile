FROM golang:1.11.5 as builder

WORKDIR /api
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -o api .

FROM debian:9.7-slim as runner

WORKDIR /api
COPY --from=builder /api/api .

EXPOSE 3000/tcp
ENTRYPOINT ["./api"]
