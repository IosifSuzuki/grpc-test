FROM golang:1.18-alpine as BuildStage

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

COPY . .

EXPOSE 8080

RUN go build -o main cmd/rest/main.go

FROM alpine:latest

WORKDIR /

COPY --from=BuildStage app/internal/config/config.yml /internal/config/config.ym
COPY --from=BuildStage app/main /main

EXPOSE 8080

CMD ["./main"]