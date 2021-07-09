FROM golang:1.16-alpine AS build

#ARG APP_PKG_NAME
ENV APP_PKG_NAME="nginx1"\
    BINARY_NAME="nginx1"\
    CGO_ENABLED=0
WORKDIR /go/src/$APP_PKG_NAME
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./vendor ./vendor

#ARG BINARY_NAME
RUN go mod init
RUN go mod vendor
RUN go vet ./...
RUN out=$(go fmt ./...) && if [[ -n "$out" ]]; then echo "$out"; exit 1; fi
RUN go test ./...

RUN go build -o /app/service ./cmd/$BINARY_NAME

CMD ["/app/service"]