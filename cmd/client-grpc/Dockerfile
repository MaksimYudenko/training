FROM golang:1.12-alpine AS build_base
EXPOSE 8080 9090 5432
ENV SRC_DIR=/github.com/MaksimYudenko/training
WORKDIR $SRC_DIR
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN apk add bash ca-certificates git gcc g++ libc-dev
RUN go mod download
FROM build_base AS builder
COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/client-grpc
FROM alpine
RUN apk add ca-certificates

COPY --from=builder /go/bin/client-grpc /bin/client-grpc
ENTRYPOINT ["./bin/client-grpc","-server=mynet:9090"]