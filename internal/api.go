package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type (
	API struct {
		URL    string
		Key    string
		Params []string
	}
)

var ErrParamArrayInvalid = errors.New("nasa-wallpaper: error, one of your parameters is missing a value")

func (api API) FullURL() string {
	return api.URL + api.evalParamString()
}

func (api API) evalParamString() string {
	if len(api.Params)%2 == 1 {
		panic(ErrParamArrayInvalid)
	}

	var params string = "?" + "api_key=" + api.Key + "&"
	for i, v := range api.Params {

		isEven := i%2 == 0

		if isEven {
			// is a parameter
			params += v + "=" + api.Params[i+1] + "&"
		} else {
			// is a value
			continue
		}
	}

	return params
}

func Get(api API) *http.Response {
	url := api.FullURL()

	fmt.Println("GET", url)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// fmt.Println("STATUS", res.Status)
	return res
}

func ProcessResponse(res *[]byte, v any) any {
	_ = json.Unmarshal(*res, &v)
	// ignoring error as it should eventually be handled properly
	return v
}
