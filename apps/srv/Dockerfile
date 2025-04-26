FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata
RUN adduser -D -g '' builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo \
    -ldflags="-w -s -X main.version=$(git describe --tags --always) -X main.buildDate=$(date +%Y-%m-%dT%H:%M:%S)" \
    -o /app/bastion ./main.go

FROM gcr.io/distroless/static:nonroot

USER nonroot:nonroot

COPY --from=builder /app/bastion /opt/bin/bastion
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

ENV TZ=Etc/UTC
ENV APP_ENV=production

ENTRYPOINT ["/opt/bin/bastion"]

CMD []
