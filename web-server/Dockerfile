FROM golang:1.19-alpine AS builder

WORKDIR /src

COPY . .

RUN go mod download 


RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

FROM alpine 

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata &&\
    apk add --no-cache go

COPY --from=builder /src/app .

ENTRYPOINT ["./main"]
