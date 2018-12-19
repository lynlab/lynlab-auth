package main

import "fmt"

type (
	ErrorOutput struct {
		StatusCode int    `json:"-"`
		Message    string `json:"message"`
	}

	TokenGenerateOutput struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

func (e *ErrorOutput) Error() string {
	return fmt.Sprintf(e.Message)
}
