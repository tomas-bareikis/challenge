FROM golang:1.21.1-alpine3.18 as builder

WORKDIR /app
ADD . /app
RUN go test main_test.go -c -o challenge
RUN chmod +x challenge
ENV LANGUAGE=go
CMD ["/app/challenge"]

FROM php:8.1-cli-alpine3.17 as php
COPY --from=builder /app/challenge /opt/challenge
COPY --from=builder /app/tests /opt/tests
COPY --from=builder /app/run-php.sh /opt/run-php.sh
VOLUME /opt/php
WORKDIR /opt
ENV LANGUAGE=php
CMD ["/opt/challenge", "-test.v"]

FROM golang:1.21.1-alpine3.18 as go
COPY --from=builder /app/challenge /opt/challenge
COPY --from=builder /app/tests /opt/tests
COPY --from=builder /app/run-go.sh /opt/run-go.sh
VOLUME /opt/go
WORKDIR /opt
ENV LANGUAGE=go
CMD ["/opt/challenge", "-test.v"]

FROM node:18.17.1-alpine3.18 as js
COPY --from=builder /app/challenge /opt/challenge
COPY --from=builder /app/tests /opt/tests
COPY --from=builder /app/run-js.sh /opt/run-js.sh
VOLUME /opt/js
WORKDIR /opt
ENV LANGUAGE=js
CMD ["/opt/challenge", "-test.v"]

FROM mcr.microsoft.com/dotnet/sdk:5.0 as cs
COPY --from=builder /app/challenge /opt/challenge
COPY --from=builder /app/tests /opt/tests
COPY --from=builder /app/run-cs.sh /opt/run-cs.sh
VOLUME /opt/cs
WORKDIR /opt
ENV LANGUAGE=cs
CMD ["/opt/challenge", "-test.v"]

FROM ubuntu:22.04 as sh
COPY --from=builder /app/challenge /opt/challenge
COPY --from=builder /app/tests /opt/tests
COPY --from=builder /app/run-sh.sh /opt/run-sh.sh
VOLUME /opt/sh
WORKDIR /opt
ENV LANGUAGE=sh
CMD ["/opt/challenge", "-test.v"]
