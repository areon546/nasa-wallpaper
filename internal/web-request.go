package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Get(api api) *http.Response {
	url := api.FullURL()

	fmt.Println("GET", url)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("STATUS", res.Status)
	return res
}

func ProcessResponse(res *http.Response, v any) any {
	dec := json.NewDecoder(res.Body)
	dec.Decode(v)
	return v
}

func GetPhoto() {}
