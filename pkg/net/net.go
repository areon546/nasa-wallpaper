package net

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/areon546/go-files/files"
)

func Get(url string) *http.Response {
	fmt.Println("GET", url)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// fmt.Println("STATUS", res.Status)
	return res
}

func ReadResponse(res *http.Response) []byte {
	contents, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	return contents
}

func ProcessResponse(res *[]byte, v any) any {
	_ = json.Unmarshal(*res, &v)
	// ignoring error as it should eventually be handled properly
	return v
}

// DownloadFile GET requests a url, and writes the entire contents
// to the given filename
func DownloadFile(filename, url string) error {
	f := files.NewFile(filename)
	_, err := f.Write(ReadResponse(Get(url)))
	return err
}
