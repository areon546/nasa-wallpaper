package nasa

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"time"

	"areon546/nasa-wallpaper/internal"
	"areon546/nasa-wallpaper/pkg/cmd"
	"areon546/nasa-wallpaper/pkg/net"
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
	something := net.ProcessResponse(res, &thing)

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

	if !internal.IsCached(apod.Date) {
		if reflect.DeepEqual(apod.MediaType, "image") {
			internal.CacheURL(apod.Date, url)
		} else {
			cmd.Notify("Failed to set background", fmt.Sprintf("APOD has type %s", apod.MediaType))
			os.Exit(1)
		}
	}

	cmd.Notify("New Background: ", apod.Title)
	cmd.SetBackground(internal.CacheFilename(apod.Date))
}

type APODS []APOD

func IsAPODS(c *internal.Config, res *[]byte) (pods APODS, ok bool) {
	var thing APODS
	something := net.ProcessResponse(res, &thing)

	pods = *something.(*APODS)
	return pods, len(pods) >= 1
}

func HandleAPODs(c *internal.Config, pods APODS) {
	rand.NewSource(time.Now().UnixNano())

	random := rand.Intn(len(pods))
	apod := pods[random]
	HandleAPOD(c, apod)
}
