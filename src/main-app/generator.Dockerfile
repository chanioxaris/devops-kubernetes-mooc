FROM golang:1.15-alpine

COPY cmd/generator src/generator
WORKDIR src/generator
RUN go build -o generator .

CMD ./generator