FROM golang:1.26 AS builder

WORKDIR /build

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /usr/bin/staticWebappServer .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /usr/bin/healthChecker healthChecker/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /usr/bin/staticWebappServer /usr/bin/
COPY --from=builder /usr/bin/healthChecker /usr/bin/

ENV PORT=8080

EXPOSE $PORT

HEALTHCHECK CMD ["/usr/bin/healthChecker"]
ENTRYPOINT ["/usr/bin/staticWebappServer"]