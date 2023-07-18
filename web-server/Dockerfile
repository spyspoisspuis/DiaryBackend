FROM golang:1.19-alpine AS builder

RUN apk add --no-cache git

WORKDIR /src/app

COPY go.mod .
COPY go.sum .

# Set the environment variables based on the build arguments
ARG SECURITY_KEY_AES
ARG SECURITY_NONCE_AES
ARG CONNECTION_DBURL
ARG CONNECTION_REDISURL
ARG CORS_TARGET

ENV SECURITY_KEY_AES=$SECURITY_KEY_AES
ENV SECURITY_NONCE_AES=$SECURITY_NONCE_AES
ENV CONNECTION_DBURL=$CONNECTION_DBURL
ENV CONNECTION_REDISURL=$CONNECTION_REDISURL
ENV CORS_TARGET=$CORS_TARGET

RUN go mod download &&\
    go mod vendor 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

FROM alpine 

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata &&\
    apk add --no-cache go

COPY --from=builder /src/app .

ENTRYPOINT ["./main"]
