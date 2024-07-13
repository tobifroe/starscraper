FROM golang:alpine

RUN apk add make

WORKDIR /app

COPY . /app

RUN make build-docker

ENTRYPOINT ["./starscraper"]