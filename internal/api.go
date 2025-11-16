package internal

import (
	"errors"
)

type (
	api struct {
		URL    string
		Key    string
		Params []string
	}
)

var ErrParamArrayInvalid = errors.New("nasa-wallpaper: error, one of your parameters is missing a value")

func (api api) FullURL() string {
	return api.URL + api.evalParamString()
}

func (api api) evalParamString() string {
	if len(api.Params)%2 == 1 {
		panic(ErrParamArrayInvalid)
	}

	var params string = "?" + "api_key=" + api.Key + "&"
	for i, v := range api.Params {

		isEven := i%2 == 0

		if isEven {
			// is a parameter
			params += v + "="
		} else {
			// is a value
			params += v + "&"
		}
	}

	return params
}
