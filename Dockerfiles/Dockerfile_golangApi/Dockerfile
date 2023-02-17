# base image
FROM golang:1.17.3-alpine3.14 as base
WORKDIR /builder
RUN apk add upx

ENV GO111MODULE=on CGO_ENABLED=0

COPY go.mod go.sum /builder/
RUN go mod download

COPY . .
# RUN mkdir -p /builder/env/
# COPY /env/.env  /builder/env/.env
COPY .env /builder/.env 
RUN go build -o /builder/main /builder/main.go
RUN upx -9 /builder/main

# runner image
FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=base /builder/main main
COPY --from=base /builder/.env .env


EXPOSE 8000
CMD ["/app/main"]