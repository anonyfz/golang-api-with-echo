package api_test

import (
	"bytes"
	"encoding/json"
	"go-echo-api/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	. "go-echo-api/api"
	"go-echo-api/validation"
)

func TestLoginAPIEmptyRequestBodyShouldReturnStatusBadRequest(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedBody := `{"status":"fail","error_message":"validate error"}`
	e := echo.New()
	e.Validator = validation.NewCustomValidator()
	request := httptest.NewRequest(http.MethodPost, "/login", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recoder := httptest.NewRecorder()
	context := e.NewContext(request, recoder)
	api := NewLoginAPI(&MockLoginService{})

	actualErr := api.Login(context)

	if assert.NoError(t, actualErr) {
		assert.Equal(t, expectedStatusCode, recoder.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(recoder.Body.String()))
	}
}

func TestLoginAPIShouldReturnStatusBadRequest(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedBody := `{"status":"fail","error_message":"validate error"}`
	e := echo.New()
	e.Validator = validation.NewCustomValidator()
	loginRequest := model.LoginRequest{}
	requestBody, _ := json.Marshal(loginRequest)
	request := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recoder := httptest.NewRecorder()
	context := e.NewContext(request, recoder)
	api := NewLoginAPI(&MockLoginService{})

	actualErr := api.Login(context)

	if assert.NoError(t, actualErr) {
		assert.Equal(t, expectedStatusCode, recoder.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(recoder.Body.String()))
	}
}

func TestLoginAPIShouldReturnStatusOK(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedBody := `{"status":"success"}`
	e := echo.New()
	e.Validator = validation.NewCustomValidator()
	loginRequest := model.LoginRequest{
		Email:    "fake@example.com",
		Password: "123456789",
	}
	requestBody, _ := json.Marshal(loginRequest)
	request := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recoder := httptest.NewRecorder()
	context := e.NewContext(request, recoder)
	api := NewLoginAPI(&MockLoginService{})

	actualErr := api.Login(context)

	if assert.NoError(t, actualErr) {
		assert.Equal(t, expectedStatusCode, recoder.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(recoder.Body.String()))
	}
}

func TestNewAPI(t *testing.T) {
	service := &MockLoginService{}
	expectedLoginAPI := &LoginAPI{
		Service: service,
	}

	actualLoginAPI := NewLoginAPI(service)

	assert.Equal(t, expectedLoginAPI, actualLoginAPI)
}
