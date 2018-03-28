package validationHelper

import (
	"net/url"
)

func Url(u string) (bool, error) {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
