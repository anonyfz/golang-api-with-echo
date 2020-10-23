package service

type (
	// LoginService is interface service
	LoginService interface {
		Login() error
	}

	loginService struct{}
)

// NewLoginService is wrap struct loginService
func NewLoginService() LoginService {
	return &loginService{}
}

// Login is handle service. It still bypass checking user
func (s *loginService) Login() error {
	return nil
}
