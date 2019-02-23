FROM golang:1:11
RUN go get gopkg.in/alecthomas/kingpin.v2
COPY data /go/src/csv2gpx/data
COPY converter.go /go/src/csv2gpx
COPY main.go /go/src/csv2gpx
WORKDIR /go/src/csv2gpx
RUN go build
RUN csv2gpx.exe data/example_input.csv data/example_output.gpx