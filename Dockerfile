FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . .
RUN go mod download
RUN GOOS=linux go build -ldflags="-s -w" -o ./server ./main.go

FROM alpine:3.10
WORKDIR /usr/bin
COPY --from=builder /go/src/app/server .
COPY --from=builder /go/src/app/data/currency.json ./data/currency.json
RUN chmod +x ./server
EXPOSE 80
ENTRYPOINT ["./server"]