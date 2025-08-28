package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {

	val := headers.Get("Authorization")
	fmt.Println("val:", val)

	if val == "" {
		return "", errors.New("no Authorization Info Found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malfunctioned headers")
	}
	if vals[0] != "api_key" {
		return "", errors.New("malfunctioned first part of the header")
	}
	return vals[1], nil
}
