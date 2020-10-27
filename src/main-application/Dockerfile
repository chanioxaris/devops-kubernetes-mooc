FROM golang:1.15-alpine

COPY cmd src/app
WORKDIR src/app
RUN go mod init && go mod download
RUN go build -o app .

EXPOSE 3000

CMD ./app