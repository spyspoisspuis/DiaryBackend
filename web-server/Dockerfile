FROM golang:1.19-alpine AS builder

RUN apk add --no-cache git

WORKDIR /tmp

COPY go.mod .
COPY go.sum .

RUN go mod download &&\
    go mod vendor

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

FROM alpine 

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata &&\
    apk add --no-cache go

COPY --from=builder /tmp .

CMD ["./main"]
