FROM golang:1.15-alpine

COPY . src/project
WORKDIR src/project
RUN go build -o project ./cmd

EXPOSE 8080

CMD ./project