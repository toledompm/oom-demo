FROM golang:1.23.0-alpine AS build

WORKDIR /app

COPY go.mod .

COPY main.go .

RUN go build -o main .

FROM alpine:3.20 AS runtime

WORKDIR /app

COPY --from=build /app/main .

ENTRYPOINT ["./main"]
