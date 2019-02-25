FROM golang:1-alpine AS builder

LABEL maintainer="Matthias Sommer, matthiassommer@posteo.de"

WORKDIR /go/src/github.com/matthiassommer/csv2gpx

# install dependencies
RUN apk add --update git
RUN go get gopkg.in/alecthomas/kingpin.v2

# copy the code and data
COPY data/example_input.csv ./data/
COPY main.go .
COPY converter.go .

# build the app
RUN GOOS=linux GOARCH=amd64 go build -o app .

# Second stage
FROM alpine

COPY --from=builder /go/src/github.com/matthiassommer/csv2gpx/ .

ENV INPUT ""
ENV OUTPUT ""

CMD ./app $INPUT $OUTPUT