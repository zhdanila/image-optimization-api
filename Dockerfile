FROM golang:1.23.4-alpine

WORKDIR /app

RUN apk update && apk add --no-cache \
    build-base \
    vips-dev \
    cairo-dev \
    pango-dev

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o main cmd/server/main.go

EXPOSE 8080

CMD [ "./main" ]
