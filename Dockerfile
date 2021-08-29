FROM golang:rc-alpine3.12 AS build

ENV GO11MODULE=on

RUN apk add build-base tzdata

WORKDIR /app

COPY . .

RUN go build

FROM alpine

COPY --from=build /app/timestamps /bin/timestamps

EXPOSE 8080

CMD ["/bin/timestamps"]
