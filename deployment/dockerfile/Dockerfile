FROM golang:1.13-alpine as builder

WORKDIR /go/src/github.com/ddddddO/extxt
COPY ./ ./

RUN go build -o extxt cmd/extxt/main.go

FROM alpine

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/ddddddO/extxt/extxt /extxt
# templateのhtmlも必要なため
COPY --from=builder /go/src/github.com/ddddddO/extxt/server/templates /server/templates
COPY ./deployment/dockerfile/secret.json /secret.json

ENV GOOGLE_APPLICATION_CREDENTIALS=/secret.json
CMD ["/extxt", "server"]
