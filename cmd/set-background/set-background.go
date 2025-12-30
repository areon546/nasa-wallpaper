package main

import (
	"fmt"
	"net/http"

	"areon546/nasa-wallpaper/internal"
	"areon546/nasa-wallpaper/pkg/nasa"
	"areon546/nasa-wallpaper/pkg/net"
)

var c *internal.Config

func main() {
	c = internal.ReadConfig()
	apiKeys := c.APIs()

	for _, api := range apiKeys {

		res := net.Get(c.Api[api].FullURL())
		if res.StatusCode != http.StatusOK {
			fmt.Println(api, res.StatusCode, http.StatusText(res.StatusCode))

			continue // skip to next one
		}

		if handle(res) {
			fmt.Println(api, "ok")
			return
		}

		fmt.Println(api, "not handled")
	}
}

// Checks pre-written handlers in order,
// and returns true if it finds it at that point.
func handle(res *http.Response) (finished bool) {
	// NOTE: since it is a response, it can only be read once.
	// I should instead be passing through the jsonBytes
	json := net.ReadResponse(res)

	if apods, ok := nasa.IsAPODS(c, &json); ok {
		nasa.HandleAPODs(c, apods)
		return true
	}
	if apod, ok := nasa.IsAPOD(c, &json); ok {
		nasa.HandleAPOD(c, apod)
		return true
	}

	fmt.Println(res.Header)

	return
}
