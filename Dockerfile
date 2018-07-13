FROM golang:alpine AS builder
RUN apk add --no-cache git
ADD . src/github.com/b00lduck/mqtt_sn_loadbalancer
WORKDIR src/github.com/b00lduck/mqtt_sn_loadbalancer
RUN go get ./... &&\
    go vet ./... &&\
    go test ./... &&\
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .

FROM scratch
USER 100:100
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
#EXPOSE 8080
