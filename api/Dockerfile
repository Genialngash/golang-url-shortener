FROM golang:alpine as builder

RUN mkdir /build

ADD ./build/

WORKDIR /build/

RUN go build -0 main .

#stage 2

FROM alpine

RUN adduser -s -p -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]

