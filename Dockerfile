FROM golang:1.19.0-bullseye AS builder

RUN go install github.com/beego/bee/v2@latest

# ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor

# ENV APP_HOME  /go/src/goapp
# RUN mkdir -p "$APP_HOME"
# WORKDIR "$APP_HOME"
# COPY src/ .

WORKDIR /app

COPY src/ ./

RUN go mod download
# go mod vendor
RUN go mod verify
RUN go build -o /goapp

FROM golang:1.19.0-bullseye
# ENV APP_HOME /go/src/goapp
# RUN mkdir -p "$APP_HOME"
# WORKDIR "$APP_HOME"
WORKDIR /


COPY src/conf/ conf/
COPY src/views/ views/
# COPY --from=builder "$APP_HOME"/goapp $APP_HOME
# COPY --from=builder $APP_HOME/goapp goapp

COPY --from=builder /goapp /goapp

EXPOSE 8080
CMD ["/goapp"]
# CMD ["./goapp"]
# CMD ["bee", "run"]