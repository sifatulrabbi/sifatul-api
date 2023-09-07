FROM golang:1.21.0-alpine as builder

WORKDIR /app

ENV GOENV=production
ENV PORT=9876
ENV SMTP_PORT_SSL=465
ENV SMTP_PORT_TLS=587

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o ./api ./main.go


FROM alpine:latest

WORKDIR /app

ENV GOENV=production
ENV PORT=9876
ENV SMTP_PORT_SSL=465
ENV SMTP_PORT_TLS=587

COPY --from=builder ./app/api ./

EXPOSE 9876

CMD [ "./api" ]
