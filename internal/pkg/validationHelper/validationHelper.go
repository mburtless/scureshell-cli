package validationHelper

import (
	"net/url"
	"os"
)

func Url(u string) (bool, error) {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func File(f string) (bool, error) {
	if _, err := os.Stat(f); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func FileExists(f string) (bool, error) {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false, err
	} else {
		return true, nil
	}
	/*if _, err := os.Stat(f); err != nil {
		return true, nil
	} else {
		return false, err
	}*/
}
