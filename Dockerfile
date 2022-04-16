FROM golang:1.14-alpine AS builder

RUN apk update && \
    apk upgrade && \
    apk add curl make \
    && mkdir -p /app 
    
WORKDIR /app
COPY    . .

EXPOSE 3000