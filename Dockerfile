FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-w -s' -o /opt/application ./src/cmd/application

FROM alpine:latest
WORKDIR /opt/
COPY --from=builder /opt/application .

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown appuser:appgroup ./application
USER appuser

ENV PORT 8181
EXPOSE $PORTT

CMD ["./application"]
