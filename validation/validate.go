package validation

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// NewCustomValidator is wrap struct customValidator
func NewCustomValidator() echo.Validator {
	return &customValidator{
		validator: validator.New(),
	}
}
