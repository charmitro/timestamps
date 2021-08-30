FROM golang:rc-alpine3.12 AS build

ENV GO11MODULE=on

RUN apk add build-base
WORKDIR /app

COPY . .

RUN go build

FROM alpine

# tzdata package needed for timezones
RUN apk add tzdata

COPY --from=build /app/timestamps /bin/timestamps

ENTRYPOINT ["/bin/timestamps"]
