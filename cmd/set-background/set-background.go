package main

import (
	"fmt"
	"log"
	"reflect"

	"areon546/nasa-wallpaper/internal"
	"areon546/nasa-wallpaper/pkg/nasa"
)

var c *internal.Config

func main() {
	c = internal.ReadConfig()
	res := internal.Get(c.Api["apod"])

	var picture nasa.APOD
	something := internal.ProcessResponse(res, &picture)

	log.Println(reflect.TypeOf(something), something)

	handle(something)
}

func handle(something any) {
	switch something := something.(type) {
	case *nasa.APOD:
		nasa.HandleAPOD(*c, something)

	case *[]nasa.APOD:
		nasa.HandleAPODs(c, *something)

	default:
		fmt.Println("Unknown API Type")

		fmt.Println(something)
	}
}
