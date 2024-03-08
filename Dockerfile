FROM golang:1.22 as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/figth_club ./cmd/main.go

FROM alpine:3.14.10

EXPOSE 8080

COPY --from=builder /app/bin/figth_club .

ENV GOGC 1000
ENV GOMAXPROCS 3

CMD ["/figth_club"]