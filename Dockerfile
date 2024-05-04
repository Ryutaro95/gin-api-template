FROM golang:1.22

ENV TZ="Asia/Tokyo"

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

RUN go get -u github.com/cosmtrek/air@latest && \
  go build -o /go/bin/air github.com/cosmtrek/air

RUN useradd -m app
RUN chown -R app:app /go/pkg

USER app

