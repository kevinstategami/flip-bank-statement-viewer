FROM golang:1.20-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://proxy.golang.org,direct

COPY . .
RUN go build -o /flip-bank-viewer ./cmd/server

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /flip-bank-viewer /flip-bank-viewer
EXPOSE 8080
ENTRYPOINT ["/flip-bank-viewer"]
