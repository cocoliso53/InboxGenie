FROM golang:1.20-alpine as build

RUN apk add --no-cache git
WORKDIR /app
COPY . .

RUN go build -o inbox_genie

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/inbox_genie .
COPY --from=build /app/config.env .

CMD ["./inbox_genie"]