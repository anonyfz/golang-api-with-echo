package api_test

type MockLoginAPI struct{}

func (MockLoginAPI) Login() error {
	return nil
}
