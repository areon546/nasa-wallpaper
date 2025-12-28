package nasa

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"time"

	"areon546/nasa-wallpaper/internal"
	"areon546/nasa-wallpaper/pkg/cmd"
)

type APOD struct {
	Copyright   string
	Date        string
	Explanation string
	Hdurl, URL  string
	MediaType   string `json:"media_type"`
	Title       string
	Version     string `json:"service_version"`
}

func IsAPOD(c *internal.Config, res *[]byte) (pod APOD, ok bool) {
	var thing APOD
	something := internal.ProcessResponse(res, &thing)

	pod = *something.(*APOD)
	return pod, len(pod.URL) >= 1
}

func HandleAPOD(c *internal.Config, apod APOD) {
	var url string

	if c.Hidef {
		url = apod.Hdurl
	} else {
		url = apod.URL
	}

	fmt.Println("URLS: ", url, c.Hidef, apod.URL, url)

	if reflect.DeepEqual(apod.MediaType, "image") {
		cmd.Notify("New Background: ", apod.Title)
		cmd.SetBackground(url)
	} else {
		cmd.Notify("Failed to set background", fmt.Sprintf("APOD has type %s", apod.MediaType))
		os.Exit(1)
	}
}

type APODS []APOD

func IsAPODS(c *internal.Config, res *[]byte) (pods APODS, ok bool) {
	var thing APODS
	something := internal.ProcessResponse(res, &thing)

	pods = *something.(*APODS)
	return pods, len(pods) >= 1
}

func HandleAPODs(c *internal.Config, pods APODS) {
	rand.NewSource(time.Now().UnixNano())

	random := rand.Intn(len(pods))
	apod := pods[random]
	HandleAPOD(c, apod)
}
