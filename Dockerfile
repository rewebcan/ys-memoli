FROM golang:alpine as builder
COPY go.mod go.sum /go/src/github.com/rewebcan/ys-memoli/
WORKDIR /go/src/github.com/rewebcan/ys-memoli/
RUN go mod download
COPY . /go/src/github.com/rewebcan/ys-memoli/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/app /go/src/github.com/rewebcan/ys-memoli/cmd/app

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
RUN mkdir -p /tmp
COPY --from=builder /go/src/github.com/rewebcan/ys-memoli/bin/app /usr/bin/app
EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/app"]