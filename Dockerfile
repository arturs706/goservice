FROM golang:1.22.2-alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64
RUN go build -ldflags="-s -w" -o gouserservice .

FROM scratch
COPY --from=builder ["/build/gouserservice", "/build/.env", "/"]

ENTRYPOINT ["/gouserservice"]