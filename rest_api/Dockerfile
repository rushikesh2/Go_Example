FROM golang:1.12.0-alpine3.9

WORKDIR /app/server
RUN apk add git
COPY . .
RUN go get ./... 
RUN go build 
CMD ["/app/main"]
