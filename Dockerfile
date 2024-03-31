FROM golang:1.20-bullseye AS builder

EXPOSE 80

WORKDIR /usr/src

COPY . .

RUN go build -o sportmen .

FROM debian:bullseye-slim

COPY --from=builder /usr/src/sportmen /usr/local/bin/sportmen

CMD ["sportmen"]