FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git 

COPY go.mod go.sum ./
RUN go mod download


COPY .  .

#build the Go binary
RUN go build -o main ./main.go


#Run Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3000

CMD [ "./main" ]

