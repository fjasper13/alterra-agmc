#build executeable binary
FROM golang:1.18-alpine AS builder

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /alterra-agmc

#build small image
FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder alterra-agmc .
#COPY .env . 
EXPOSE 8080

CMD [ "./alterra-agmc" ]