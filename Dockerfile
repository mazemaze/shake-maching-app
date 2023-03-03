FROM golang:1.17-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
ENV GO111MODULE=on
WORKDIR ${ROOT}

RUN apk update && apk add git
RUN go get github.com/cosmtrek/air
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
