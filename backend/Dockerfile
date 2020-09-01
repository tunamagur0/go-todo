FROM golang:1.14-alpine as builder

RUN mkdir /go_todo/
ADD ./ /go_todo/
WORKDIR /go_todo/src

RUN apk update && apk add --update sqlite libc-dev musl-dev gcc
RUN go build -o main ./

FROM alpine:latest

RUN apk update && apk add --update sqlite libc-dev musl-dev gcc

WORKDIR /app/
COPY --from=builder /go_todo/src/main /app/
RUN chmod +x main

CMD ./main $PORT

