package main

import "fmt"

type (
	ErrorOutput struct {
		StatusCode int    `json:"-"`
		Message    string `json:"message"`
	}
)

func (e *ErrorOutput) Error() string {
	return fmt.Sprintf(e.Message)
}
