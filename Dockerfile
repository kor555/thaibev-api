
FROM golang:1.27.2-alpine

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

ENV PATH="/go/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download

CMD ["CompileDaemon", "-directory=/app", "-build=go build -o /tmp/app ./cmd/main.go", "-command=/tmp/app"]
