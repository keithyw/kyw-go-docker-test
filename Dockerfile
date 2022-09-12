FROM golang:1.19.0-bullseye AS builder

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /app

COPY src/ ./

RUN go mod download
RUN go mod verify
RUN go build -o /goapp

FROM golang:1.19.0-bullseye

WORKDIR /

COPY src/conf/ conf/
COPY src/views/ views/
COPY --from=builder /goapp /goapp

EXPOSE 8080
CMD ["/goapp"]