# Build Stage
FROM golang:1.20rc3-alpine3.17 as builder
WORKDIR /app
COPY . .
# RUN go build -o main cmd/main.go

# Run Stage
# FROM alpine:3.16

RUN apk update &&  apk add git
RUN go install github.com/cosmtrek/air@latest

# WORKDIR /app
# COPY --from=builder /app/main .
COPY development.env .
# COPY start.sh .
# COPY wait-for.sh .
# COPY db/migration ./db/migration

RUN go mod tidy

EXPOSE 8080
# CMD [ "/app/main" ]
CMD ["air"]
# ENTRYPOINT [ "/app/start.sh" ]
