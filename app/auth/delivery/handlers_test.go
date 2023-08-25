package authdelivery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	authinterfaces "github.com/masharpik/TransactionalSystem/app/auth/interfaces"
	"github.com/masharpik/TransactionalSystem/app/auth/utils"
)

func TestAuthHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := authinterfaces.NewMockIAuthService(mockCtrl)

	router := &Delivery{
		service: mockService,
	}

	expectedUser := utils.User{UserID: "73a94f3a-42cb-11ee-9b4a-0242c0a81003", Balance: 0}
	mockService.EXPECT().CreateUser().Return(expectedUser, nil)

	req, err := http.NewRequest("POST", "api/auth", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router.AuthHandler(rr, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, rr.Code, "Unexpected HTTP status code")
}
