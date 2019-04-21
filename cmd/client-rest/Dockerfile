FROM golang:1.12-alpine AS build_base

EXPOSE 8080
EXPOSE 5432
ENV SRC_DIR=/go/src/github.com/MaksimYudenko/training/cmd/client-rest
# ENV GO111MODULE=on
WORKDIR $SRC_DIR

RUN apk add bash ca-certificates git gcc g++ libc-dev

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

FROM build_base AS client_builder
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./

FROM alpine
RUN apk add ca-certificates

COPY --from=client_builder /go/bin/client-rest /bin/client-rest

ENTRYPOINT ["./bin/client-rest","-server=http://localhost:8080"]