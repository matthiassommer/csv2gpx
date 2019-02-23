package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inputPath  = kingpin.Arg("input", "Path to input file.").Required().String()
	outputPath = kingpin.Arg("output", "Path to output file.").Required().String()
)

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()

	Convert(*inputPath, *outputPath)
}
