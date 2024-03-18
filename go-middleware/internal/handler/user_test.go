package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/model"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/service/mocks"
)

func TestUserSignUp(t *testing.T) {
	t.Run("error binding", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		req := httptest.NewRequest(http.MethodPost, "/users/sign-up", bytes.NewBuffer([]byte(`{"username":"","password":""}`)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(rec)
		g.Request = req

		usrHdl := userHandlerImpl{}
		usrHdl.UserSignUp(g)

		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
	t.Run("error binding password", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		req := httptest.NewRequest(http.MethodPost, "/users/sign-up", bytes.NewBuffer([]byte(`{"username":"username","password":"abc"}`)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		// gin context mock
		g, _ := gin.CreateTestContext(rec)
		g.Request = req

		usrHdl := userHandlerImpl{}
		usrHdl.UserSignUp(g)

		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})

	t.Run("error sign up service", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		req := httptest.NewRequest(http.MethodPost, "/users/sign-up", bytes.NewBuffer([]byte(`{"username":"username","password":"abc12345"}`)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		// gin context mock
		g, _ := gin.CreateTestContext(rec)
		g.Request = req

		svcMock := mocks.NewUserService(t)
		svcMock.
			On("SignUp", g, model.UserSignUp{Username: "username", Password: "abc12345"}).
			Return(model.User{}, errors.New("some error"))

		usrHdl := userHandlerImpl{svc: svcMock}
		usrHdl.UserSignUp(g)

		assert.Equal(t, http.StatusInternalServerError, rec.Result().StatusCode)
	})
}
