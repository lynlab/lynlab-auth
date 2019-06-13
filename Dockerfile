### Builder
FROM golang:1.12-alpine as builder

RUN apk update && apk add git && apk add ca-certificates

WORKDIR /usr/src/app
ENV GO11MODULE on

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .


### Make executable image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/src/app/main /main
COPY --from=builder /usr/src/app/public /public

ENTRYPOINT [ "/main" ]
