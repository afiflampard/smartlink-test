FROM golang:alpine3.13 as builder

RUN apk update && apk upgrade && \
  apk --no-cache --update add git make

WORKDIR /app

COPY . .

RUN go mod download && \
  go build -v -o engine && \
  chmod +x engine

## Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --no-cache --update add ca-certificates tzdata && \
  mkdir /app

WORKDIR /app

EXPOSE 8000

COPY --from=builder /app /app

CMD /app/engine