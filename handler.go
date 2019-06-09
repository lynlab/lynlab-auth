package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type signinInput struct {
	AppUUID  string `json:"appId"`
	Provider string `json:"provider"`
	Payload  string `json:"payload"`
}

type signinOutput struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireAt     int64  `json:"expireAt"`
}

func signin(c echo.Context) error {
	var input signinInput
	c.Bind(&input)

	switch input.Provider {
	case "google":
		u, err := getGoogleUser(input.Payload)
		if err != nil {
			return errorAPI(c, http.StatusUnauthorized, "unauthorized")
		}

		var account UserAccount
		db.Where(&UserAccount{Provider: "google", ProviderIdentity: u.Email}).First(&account)
		if account.ID == 0 {
			return errorAPI(c, http.StatusUnauthorized, "no_such_account")
		}

		// TODO - check if identity with email already exists

		identity := account.GetIdentity()

		var app Application
		db.Where(&Application{UUID: input.AppUUID}).First(&app)

		var allowed int
		db.Where(&UserAllowedApplication{ApplicationID: app.ID, IdentityID: identity.ID}).Count(&allowed)
		if allowed == 0 {
			return errorAPI(c, http.StatusUnauthorized, "authorization_required")
		}

		token := identity.NewToken(app.Scopes)
		output := signinOutput{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			ExpireAt:     token.AccessTokenExpireAt.Unix() * 1000,
		}
		return c.JSON(http.StatusOK, output)

	default:
		return errorAPI(c, http.StatusBadRequest, "no_such_provider")
	}
}
