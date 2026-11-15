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

FROM gcr.io/distroless/base-debian13:nonroot AS runner
RUN apt-get update && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/checkers-backend /app
EXPOSE 9876
USER nonroot
ENTRYPOINT ["./checkers-backend"]