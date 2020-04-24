package main

import (
	"fmt"
	"repos/week3"
	"repos/week3/exporter"
	facebook "repos/week3/facebook"
	"repos/week3/linkedin"
	"repos/week3/twitter"
)

func main() {
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdin := new(linkedin.Linkedin)
	err := exporter.Export(twtr, "twtrdata.txt")
	err = exporter.Export(fb, "fbdata.txt")
	err = exporter.Export(lnkdin, "lnkdin.txt")

	err = exporter.JSONExporter(fb, "facebook.json")
	err = exporter.JSONExporter(twtr, "twitter.json")
	err = exporter.JSONExporter(lnkdin, "linkedin.json")

	err = exporter.XMLExporter(fb, "facebook.xml")
	err = exporter.XMLExporter(twtr, "twitter.xml")
	err = exporter.XMLExporter(lnkdin, "linkedin.xml")

	if err != nil {
		panic(err)
	}
}

// ScrollFeeds prints all social media feeds
func ScrollFeeds(platforms ...week3.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("=================================")
	}
}
