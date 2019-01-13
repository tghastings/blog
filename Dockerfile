FROM golang:1.11.4 as builder
ADD . /go/src/blog
WORKDIR /go/src/blog
RUN go get blog
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /go/src/blog/main /app/
WORKDIR /app
CMD ["./main"]

EXPOSE 8090