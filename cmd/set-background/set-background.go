package main

import (
	"fmt"
	"os/exec"

	"areon546/nasa-wallpaper/internal"
)

var c *internal.Config

type APOD struct {
	Copyright      string
	Date           string
	Explanation    string
	Hdurl, URL     string
	MediaType      string
	Title          string
	ServiceVersion string
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

	fmt.Println("URLS: ", url, c.Hidef, apod.URL)

	// alternative bg args : see man feh /--bg-
	// --bg-scale
	// --bg-max
	// --bg-tile
	// --bg-fill
	cmd := exec.Command("/usr/bin/feh", "--bg-max", url)

	fmt.Println("Command ran: \n", cmd.String())

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
