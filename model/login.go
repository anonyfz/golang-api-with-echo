package model

type (
	// LoginRequest is model for login request body
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	// LoginResponse is model for login response body
	LoginResponse struct {
		Status string `json:"status"`
	}
)
