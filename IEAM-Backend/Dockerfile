FROM golang:1.17-alpine3.14
# RUN apk add build-base
WORKDIR /app

COPY . .
# RUN go mod download
RUN go build 
EXPOSE 8080

ENTRYPOINT ["./ieam-backend"]