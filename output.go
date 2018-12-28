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

	MeOutput struct {
		UUID     string `json:"uuid"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}
)

func (e *ErrorOutput) Error() string {
	return fmt.Sprintf(e.Message)
}
