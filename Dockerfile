FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY ./vendor ./vendor
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -o /jumpshot-ext-authz github.com/ramonberrutti/jumpshot-ext-authz/cmd/jumpshot-ext-authz

FROM alpine:3.19

WORKDIR /

COPY --from=build /jumpshot-ext-authz /jumpshot-ext-authz

EXPOSE 9090

ENTRYPOINT ["/jumpshot-ext-authz"]
