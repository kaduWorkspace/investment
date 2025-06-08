FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0  \
    GOARCH="amd64" \
    GOOS=linux

WORKDIR /build
COPY . .

COPY go.mod .
COPY go.sum .
RUN go mod tidy
RUN go build --ldflags "-extldflags -static" -o main cmd/web/main.go

FROM alpine:latest
WORKDIR /www

COPY --from=builder /build/main /www/
COPY --from=builder /build/core/interfaces/web/ /www/core/interfaces/web/
COPY --from=builder /build/public/ /www/public/
COPY --from=builder /build/resources/ /www/resources/
COPY --from=builder /build/.env /www/.env

ENTRYPOINT ["/www/main"]
