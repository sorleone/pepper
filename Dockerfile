FROM golang:1.14.6-alpine3.12 as builder

RUN apk add --no-cache build-base=0.5-r2

WORKDIR /go/src
COPY src .
RUN go test -v -cover ./... \
    && env CGO_ENABLED=0 go build -ldflags '-s' -o ../bin/pepper

FROM scratch

COPY --from=builder /go/bin/pepper .

ENTRYPOINT [ "./pepper" ]
