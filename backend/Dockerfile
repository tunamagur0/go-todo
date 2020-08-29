FROM golang:1.14-alpine as builder

RUN mkdir /go_todo/
WORKDIR /go_todo/
ADD . ./

RUN go build -o main ./src

FROM alpine:latest

RUN apk update && apk add sqlite

WORKDIR /app/
COPY --from=builder /go_todo/main /app/
RUN chmod +x main

CMD ["./main"]

