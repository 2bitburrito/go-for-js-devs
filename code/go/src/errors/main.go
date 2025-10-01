package main

import (
	"errors"
	"fmt"
)

type HTTPError struct {
	Code int
	Msg  string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("http %d: %s", e.Code, e.Msg)
}

func GetProfile(id string) (string, error) {
	if id != "123" {
		return "", &HTTPError{
			Code: 404,
			Msg:  "profile not found",
		}
	}
	return "Kaladin Stormblessed", nil
}

func main() {
	profile, err := GetProfile("999")
	if err != nil {
		var httpErr *HTTPError
		if errors.As(err, &httpErr) {
			fmt.Printf("got HTTP error: %d\n", httpErr.Code)
			return
		}
	}
	fmt.Println("Found Profile: ", profile)
}
