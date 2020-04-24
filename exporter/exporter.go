package exporter

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"repos/week3"
)

//Export Function
func Export(u week3.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}

//JSONExporter Function
func JSONExporter(platforms week3.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return errors.New("An error occured opening the file: " + err.Error())
	}

	feed := platforms.Feed()
	feedmapjs := make(map[int]interface{})
	for n, feedjs := range feed {
		n++
		feedmapjs[n] = feedjs
	}

	jsonfeed, err := json.Marshal(feedmapjs)
	if err != nil {
		return errors.New("an error occured during encoding: " + err.Error())
	}

	n, err := f.Write([]byte(string(jsonfeed) + "\n"))
	if err != nil {
		return errors.New("an error occured writing to file: " + err.Error())
	}
	fmt.Printf("wrote %d bytes\n", n)
	return nil
}

//XMLExporter Function
func XMLExporter(platforms week3.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return errors.New("An error occured opening the file: " + err.Error())
	}

	feed := platforms.Feed()
	xmlfeed, err := xml.MarshalIndent(feed, " ", "  ")
	if err != nil {
		return errors.New("an error occured during encoding: " + err.Error())
	}

	n, err := f.Write([]byte(string(xmlfeed) + "\n" + "\n"))
	if err != nil {
		return errors.New("An error occured writing to file" + err.Error())
	}
	fmt.Printf("Wrote %d bytes\n", n)
	return nil
}
