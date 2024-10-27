FROM golang:alpine

RUN apk --no-cache add make

WORKDIR /app

COPY . /app

RUN make build-docker

ENTRYPOINT ["./starscraper"]
