FROM golang:1.15-alpine

COPY cmd/server src/server
WORKDIR src/server
RUN go build -o server .

EXPOSE 3000

CMD ./server