# syntax=docker/dockerfile:1

FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./

# IF LOCATED IN MAINLAND CHINA, UNCOMMENT THESE 2 LINES BELOW
#RUN go env -w GO111MODULE=on
#RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go mod download -x
COPY . ./
RUN go version
RUN CGO_ENABLED=0 go build --ldflags="-s -w" -o checkers-backend

FROM alpine:latest AS runner
RUN apk update && apk upgrade --no-cache
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/checkers-backend /app
EXPOSE 9876
USER 10001:10001
ENTRYPOINT ["./checkers-backend"]