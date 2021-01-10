package main

import (
	"flag"
	"os"

	"github.com/gookit/color"
)

func main() {

	webstreamUrl := flag.String("url", "", "the url of the webstream")
	//minimumSilence := flag.Int("minsilence", 120, "minimum time to elapse until silence will be detected")
	//maxmimumLoudness := flag.Int("maxloudness", -20, "maximum loudness which will still be considered as silence")
	//checkInterval := flag.Int("checkinterval", 1000, "interval the silence should be checked [ms]")

	flag.Parse()

	if *webstreamUrl == "" { // webstream url is mandatory. if it is missing we stop and print the usage!
		color.Red.Println("The webstreamURL is mandatory!")
		flag.Usage()
		os.Exit(2)
	}

}
