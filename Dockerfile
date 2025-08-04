FROM golang:1.24.5 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o drycharting main.go

FROM alpine:3.22.0

WORKDIR /app
COPY --from=build /app/drycharting .
EXPOSE 8080

CMD ["./drycharting"]
