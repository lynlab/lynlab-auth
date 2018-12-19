package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	recoveryResetInputText = `
	{
	  "recovery_token": "example_token",
	  "password": "example_password"
	}
	`
	recoverySendEmailInputText = `
	{
	  "email": "example@example.com"
	}
	`
	registerInputText = `
	{
	  "username": "example_username",
	  "email": "example@example.com",
	  "password": "example_password"
	}
	`
	tokenGenerateInputText = `
	{
	  "email": "example@example.com",
	  "password": "example_password"
	}
	`
	tokenRefreshInputText = `
	{
	  "refresh_token": "example_token"
	}
	`
)

func TestInput(t *testing.T) {
	var i1 recoveryResetInput
	json.Unmarshal([]byte(recoveryResetInputText), &i1)
	if i1.Password != "example_password" || i1.RecoveryToken != "example_token" {
		fmt.Println("Failed to unmarshal recoveryResetInputText")
		t.Fail()
	}

	var i2 recoverySendEmailInput
	json.Unmarshal([]byte(recoverySendEmailInputText), &i2)
	if i2.Email != "example@example.com" {
		fmt.Println("Failed to unmarshal recoverySendEmailInputText")
		t.Fail()
	}

	var i3 RegisterInput
	json.Unmarshal([]byte(registerInputText), &i3)
	if i3.Username != "example_username" || i3.Email != "example@example.com" || i3.Password != "example_password" {
		fmt.Println("Failed to unmarshal registerInputText")
		t.Fail()
	}

	var i4 TokenGenerateInput
	json.Unmarshal([]byte(tokenGenerateInputText), &i4)
	if i4.Email != "example@example.com" || i4.Password != "example_password" {
		fmt.Println("Failed to unmarshal tokenGenerateInputText")
		t.Fail()
	}

	var i5 tokenRefreshInput
	json.Unmarshal([]byte(tokenRefreshInputText), &i5)
	if i5.RefreshToken != "example_token" {
		fmt.Println("Failed to unmarshal tokenRefreshInputText")
		t.Fail()
	}
}
