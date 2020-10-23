package service

type (
	LoginService interface {
		Login() error
	}

	loginService struct{}
)

func NewLoginService() LoginService {
	return &loginService{}
}

// Login is handle service. It still bypass checking user
func (s *loginService) Login() error {
	return nil
}
