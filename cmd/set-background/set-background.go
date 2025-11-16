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

	var picture any

	fmt.Println("p", &picture)

	something := internal.ProcessResponse(res, &picture)

	fmt.Println(something)
	switch something := something.(type) {
	case *APOD:
		handleAPOD(something)

	case *int: // ptr
		fmt.Println("Pointer")
	default:
		fmt.Println("Unknown API Type")

		fmt.Println(something)
	}
}

func handleAPOD(apod *APOD) {
	var url string

	fmt.Println(apod)

	if c.Hidef {
		url = apod.Hdurl
	} else {
		url = apod.URL
	}

	fmt.Println(url, c.Hidef, apod.URL)

	cmd := exec.Command("/usr/bin/feh", "--bg-max", url)

	fmt.Println(cmd.String())

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
