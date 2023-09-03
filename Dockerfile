FROM golang:1.21.0 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o ./sifatul-api ./main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/sifatul-api ./

# Set environment variables with build arguments
ENV GOENV=production
ENV PORT=9876
ENV SMTP_PORT_SSL=465
ENV SMTP_PORT_TLS=587
ENV SMTP_HOST=
ENV EMAIL_ACCOUNT=
ENV EMAIL_PASSWORD=

EXPOSE $PORT

CMD [ "./sifatul-api" ]
