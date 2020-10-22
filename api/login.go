package api

import (
	"go-echo-api/service"
)

// LoginAPI is struct api
type LoginAPI struct {
	Service service.LoginService
}

// NewLoginAPI is wrap struct *loginAPI
func NewLoginAPI(service service.LoginService) *LoginAPI {
	return &LoginAPI{
		Service: service,
	}
}
