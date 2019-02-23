package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	// GC Code must start with gc, followed by a group of 2 to 6 letters and numbers
	regExpGcCode = regexp.MustCompile("(GC|gc|Gc|gC)[a-zA-Z0-9]{2,6}$")
	regExpLat    = regexp.MustCompile("[NS°]")
	regExpLon    = regexp.MustCompile("[EW°]")
)

//Convert the input csv file to gpx
func Convert(inFile string, outFile string) {
	f, err := os.Open(inFile)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	r.Comma = ';'
	r.Comment = '#'

	output := "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<gpx xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" version=\"1.0\" creator=\"Groundspeak Pocket Query\" xsi:schemaLocation=\"http://www.topografix.com/GPX/1/0 http://www.topografix.com/GPX/1/0/gpx.xsd http://www.groundspeak.com/cache/1/0/1 http://www.groundspeak.com/cache/1/0/1/cache.xsd\" xmlns=\"http://www.topografix.com/GPX/1/0\">\n"

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		gccode := parseGcCode(row[0])

		lat := row[1]
		roundedLat := parseCoordinate(lat, "S", regExpLat)

		lon := row[2]
		roundedLon := parseCoordinate(lon, "W", regExpLon)

		// build row and append to gpx, e.g. <wpt lat=50.565150 lon=-113.716433><name>GC80AD</name></wpt>
		convertedRow := "<wpt lat=" + roundedLat + " lon=" + roundedLon + "><name>" + gccode + "</name></wpt>\n"
		output += convertedRow
	}

	output += "</gpx>"

	writeGpx(outFile, output)

}

func parseGcCode(data string) string {
	gccode := strings.TrimSpace(data)
	validGcCode := regExpGcCode.MatchString(gccode)
	if !validGcCode {
		fmt.Println("Invalid GC Code: ", gccode)
		os.Exit(1)
	}

	return gccode
}

func parseCoordinate(value string, letter string, regexp *regexp.Regexp) string {
	// get minutes and degree, e.g. N50° 33.909
	split := regexp.Split(value, -1)

	// degree is 50
	deg, _ := strconv.ParseFloat(strings.TrimSpace(split[1]), 32)

	// minutes is 33.909
	min, _ := strconv.ParseFloat(strings.TrimSpace(split[2]), 32)

	// convert decimal minutes to decimal degrees: degree + minutes/60, e.g. 50+(33.909/60)
	converted := deg + min/60

	if strings.Contains(value, letter) {
		converted *= -1
	}

	// round
	return fmt.Sprintf("%.6f", converted)
}

func writeGpx(target string, data string) {
	f, err := os.Create(target)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		panic(err)
	}
}
