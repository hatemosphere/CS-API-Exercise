FROM golang:1.11.5 as builder

WORKDIR /api

COPY . ./

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -o api .


FROM alpine:3.8 as runner
# RUN apk --no-cache add ca-certificates

WORKDIR /api
COPY --from=builder /api/api .

EXPOSE 3000/tcp
ENTRYPOINT ["./api"]
