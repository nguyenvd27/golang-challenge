# syntax=docker/dockerfile:1

FROM golang:1.17.1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./main

EXPOSE 8001

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# CMD [ "./main" ]
