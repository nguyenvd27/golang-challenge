# syntax=docker/dockerfile:1

FROM golang:1.17.1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
# COPY ./entrypoint.sh /entrypoint.sh

RUN go build -o /main

EXPOSE 8001

CMD [ "/main" ]
# CMD ["sh", "-c", "USER_NAME=root PASSWORD= /main"]
