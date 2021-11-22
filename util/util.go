package util

import (
	"errors"
	"net/url"
)

func IsURL(str string) error {
	if str == "" {
		return errors.New("url empty string")
	}
	_, err := url.Parse(str)
	return err
}
