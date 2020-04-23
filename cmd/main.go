package main

import (
	"errors"
	"fmt"
	"os"
	"repos/week3"
	"repos/week3/facebook"
	"repos/week3/twitter"
)

func main() {
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	err := export(twtr, "twtrdata.txt")
	err = export(fb, "fbdata.txt")

	if err != nil {
		panic(err)
	}

	// twtr := new(twitter.Twitter)
	// lnkdin := new(linkedin.Linkedin)

	// ScrollFeeds(fb, twtr, lnkdin)
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

func export(u week3.SocialMedia, filename string) error {
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

func exportjson(u week3.SocialMedia, filename string) error {
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
