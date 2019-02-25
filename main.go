package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inputPath  = kingpin.Arg("input", "Path to input file.").Required().String()
	outputPath = kingpin.Arg("output", "Path to output file.").Required().String()
)

func main() {
	kingpin.Version("1.0.0")
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate)

	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')

	kingpin.Parse()

	fmt.Println("CSV:", *inputPath, "GPX:", *outputPath)
	Convert(*inputPath, *outputPath)
}
