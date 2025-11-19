package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"

	"areon546/nasa-wallpaper/internal"
)

var c *internal.Config

type APOD struct {
	Copyright   string
	Date        string
	Explanation string
	Hdurl, URL  string
	MediaType   string `json:"media_type"`
	Title       string
	Version     string `json:"service_version"`
}

func main() {
	c = internal.ReadConfig()
	res := internal.Get(c.Api["apod"])

	var picture APOD
	something := internal.ProcessResponse(res, &picture)

	switch something := something.(type) {
	case *APOD:
		handleAPOD(something)

	default:
		fmt.Println("Unknown API Type")

		fmt.Println(something)
	}
}

func handleAPOD(apod *APOD) {
	var url string

	if c.Hidef {
		url = apod.Hdurl
	} else {
		url = apod.URL
	}

	fmt.Println("URLS: ", url, c.Hidef, apod.URL, url)

	if reflect.DeepEqual(apod.MediaType, "image") {
		notify("New Background: ", apod.Title)
		setBackground(url)
	} else {
		notify("Failed to set background", fmt.Sprintf("APOD has type %s", apod.MediaType))
		os.Exit(1)
	}
}

func notify(args ...string) {
	runCommand("notify-send", args...)
}

func setBackground(url string) {
	// alternative bg args : see man feh /--bg-
	// --bg-scale
	// --bg-max
	// --bg-tile
	// --bg-fill
	runCommand("/usr/bin/feh", "--bg-max", url)
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)

	fmt.Println("Command ran: \n", cmd.String())

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
