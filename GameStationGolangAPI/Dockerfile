

#  Build stage 
FROM golang:latest AS builder 
WORKDIR /app 
COPY . . 
RUN export GO111MODULE=on 
RUN go build -o main main.go 

# Run stage 
FROM alpine:latest
RUN export GO111MODULE=on 
WORKDIR /app 
COPY --from=builder /app/main .  

# this shoud not be run in production
# COPY .env . 

EXPOSE 8080 
CMD ["/app/main"]