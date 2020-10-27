FROM golang:1.15-alpine

COPY . src/ping-pong
WORKDIR src/ping-pong
RUN go build -o pingpong ./cmd

EXPOSE 3000

CMD ./pingpong