package main

type (
	recoveryResetInput struct {
		RecoveryToken string `json:"recovery_token"`
		Password      string `json:"password"`
	}

	recoverySendEmailInput struct {
		Email string `json:"email"`
	}

	RegisterInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	TokenGenerateInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	tokenRefreshInput struct {
		RefreshToken string `json:"refresh_token"`
	}
)
