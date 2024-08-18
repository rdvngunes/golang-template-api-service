# base go image
FROM golang:1.18-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o golang-template-api-service .
RUN chmod +x /app/golang-template-api-service
#build a tiny docker image
FROM alpine:latest
#Setting ENV as dev this need be ARG if need to change Run time
ENV APP_ENV=dev
RUN mkdir /app
RUN mkdir /docs
RUN mkdir /logs
#Moving Configs
COPY --from=builder /app/app/config/ /app/config
#Moving Swagger Docs
COPY --from=builder /app/docs/ /docs
#Moving build file
COPY --from=builder /app/golang-template-api-service /app
RUN apk --no-cache add chromium
EXPOSE 3000
CMD ["/app/golang-template-api-service"]