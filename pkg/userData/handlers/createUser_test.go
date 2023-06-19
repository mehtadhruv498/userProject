package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"example/userproject/mocks"
	Model "example/userproject/pkg/userData/Models"
	"example/userproject/pkg/userData/handlers"
	"fmt"

	//"example/userproject/pkg/userData/handlers"
	//"example/userproject/pkg/userData/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignUp(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockCreateUserService := new(mocks.Service)

		mockCreateUserService.On("CreateUser", mock.Anything).Return(nil)

		gin := gin.Default()

		rec := httptest.NewRecorder()

		pc := &handlers.Handler{
			Service: mockCreateUserService,
		}
		fmt.Println(pc)
		gin.POST("/api/v1/user", pc.CreateUser)

		body := Model.User{
			ID:      15,
			Name:    "test",
			Gender:  "test@gmail.com",
			DOB:     "1000",
			Address: "mathikere",
			Country: "india",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(body)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/user", &buf)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		mockCreateUserService.AssertExpectations(t)
	})
	t.Run("failure", func(t *testing.T) {
		mockSignUpService := new(mocks.Service)

		mockSignUpService.On("CreateUser", mock.Anything).Return(errors.New("some error"))

		gin := gin.Default()

		rec := httptest.NewRecorder()

		pc := handlers.Handler{
			Service: mockSignUpService,
		}
		fmt.Println(pc)
		gin.POST("/api/v1/user", pc.CreateUser)

		body := Model.User{
			ID:      15,
			Name:    "test",
			Gender:  "test@gmail.com",
			DOB:     "1000",
			Address: "mathikere",
			Country: "india",
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(body)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/user", &buf)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		mockSignUpService.AssertExpectations(t)
	})
}
