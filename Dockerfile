# ---------- BUILD STAGE ----------
FROM golang:1.20-alpine AS build

WORKDIR /app

# Copy go.mod first for caching
COPY go.mod ./
RUN go mod download

# Copy the rest of project
COPY . .

# Build binary
RUN go build -o flip-bank-viewer ./cmd/server

# ---------- RUN STAGE ----------
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/flip-bank-viewer .

EXPOSE 8080

ENTRYPOINT ["./flip-bank-viewer"]
