
FROM golang:1.22.3-alpine AS builder


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .


FROM alpine:latest
WORKDIR /root/

RUN apk add --no-cache openssl
ENV DOCKERIZE_VERSION v0.7.0
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz


RUN mkdir -p logs

COPY --from=builder /app/main .

EXPOSE 7080

CMD ["sh", "-c", "dockerize -wait tcp://mysql:3306 -wait tcp://rabbitmq:5672 -timeout 60s ./main 2>&1 | tee -a logs/app.log"]
