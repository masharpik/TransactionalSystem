package transactiondelivery

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mailru/easyjson"
	"github.com/stretchr/testify/assert"

	authUtils "github.com/masharpik/TransactionalSystem/app/auth/utils"
	transactioninterfaces "github.com/masharpik/TransactionalSystem/app/transaction/interfaces"
	"github.com/masharpik/TransactionalSystem/app/transaction/utils"
)

func TestInputHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := transactioninterfaces.NewMockITransactionService(mockCtrl)

	router := &Delivery{
		service: mockService,
	}

	testTransaction := utils.InputTransaction{UserID: "73a94f3a-42cb-11ee-9b4a-0242c0a81003", Amount: 10}
	expectedUser := authUtils.User{UserID: "73a94f3a-42cb-11ee-9b4a-0242c0a81003", Balance: 10}
	mockService.EXPECT().InputMoney(testTransaction.UserID, testTransaction.Amount).Return(expectedUser, nil)

	payload, err := easyjson.Marshal(testTransaction)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/api/transaction/input", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.InputHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Unexpected HTTP status code")
}

func TestOutputHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := transactioninterfaces.NewMockITransactionService(mockCtrl)

	router := &Delivery{
		service: mockService,
	}

	testTransaction := utils.OutputTransaction{UserID: "73a94f3a-42cb-11ee-9b4a-0242c0a81003", Amount: 10, Link: "http://example.com"}
	expectedStatus := utils.StatusTransaction{UserID: "73a94f3a-42cb-11ee-9b4a-0242c0a81003", Status: fmt.Sprintf("Информация по результату снятия придет по ссылке: %s", testTransaction.Link)}
	mockService.EXPECT().OutputMoney(testTransaction.UserID, testTransaction.Amount, testTransaction.Link).Return(expectedStatus, nil)

	payload, err := easyjson.Marshal(testTransaction)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/api/transaction/output", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.OutputHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Unexpected HTTP status code")
}

func TestInputHandler_InvalidAmounts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := transactioninterfaces.NewMockITransactionService(mockCtrl)

	router := &Delivery{
		service: mockService,
	}

	testCases := []struct {
		amount         float64
		expectedStatus int
		expectedMsg    string
	}{
		{-10, http.StatusUnprocessableEntity, utils.LogSignInputNotCorrectlyError},
		{10.999, http.StatusUnprocessableEntity, utils.LogLengthInputMoneyNotCorrectlyError},
	}

	for _, tc := range testCases {
		testTransaction := utils.InputTransaction{UserID: "73a94f3a-42cb-11ee-9b4a-0242c0a81003", Amount: tc.amount}
		payload, err := easyjson.Marshal(testTransaction)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("PUT", "/api/transaction/input", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router.InputHandler(rr, req)

		assert.Equal(t, tc.expectedStatus, rr.Code, "Unexpected HTTP status code for amount: ", tc.amount)
	}
}

func TestOutputHandler_InvalidAmounts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := transactioninterfaces.NewMockITransactionService(mockCtrl)

	router := &Delivery{
		service: mockService,
	}

	testCases := []struct {
		amount         float64
		expectedStatus int
		expectedMsg    string
	}{
		{-10, http.StatusUnprocessableEntity, utils.LogSignInputNotCorrectlyError},
		{10.999, http.StatusUnprocessableEntity, utils.LogLengthInputMoneyNotCorrectlyError},
	}

	for _, tc := range testCases {
		testTransaction := utils.OutputTransaction{UserID: "73a94f3a-42cb-11ee-9b4a-0242c0a81003", Amount: tc.amount}
		payload, err := easyjson.Marshal(testTransaction)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("PUT", "/api/transaction/output", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router.OutputHandler(rr, req)

		assert.Equal(t, tc.expectedStatus, rr.Code, "Unexpected HTTP status code for amount: ", tc.amount)
	}
}

func TestInputHandler_ServiceErrors(t *testing.T) {
    testCases := []struct {
        name           string
        errorToReturn  error
        expectedStatus int
    }{
        {
            name:           "Service returns user not found error",
            errorToReturn:  fmt.Errorf(authUtils.LogUserNotFoundError),
            expectedStatus: http.StatusUnauthorized,
        },
        {
            name:           "Service returns some other error",
            errorToReturn:  fmt.Errorf("Some"),
            expectedStatus: http.StatusInternalServerError,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            mockCtrl := gomock.NewController(t)
            defer mockCtrl.Finish()

            mockService := transactioninterfaces.NewMockITransactionService(mockCtrl)

            router := &Delivery{
                service: mockService,
            }

            testTransaction := utils.InputTransaction{UserID: "73a94f3a-42cb-11ee-9b4a-0242c0a81003", Amount: 10}
            var dummyUser authUtils.User
            mockService.EXPECT().InputMoney(testTransaction.UserID, testTransaction.Amount).Return(dummyUser, tc.errorToReturn)

            payload, err := easyjson.Marshal(testTransaction)
            if err != nil {
                t.Fatal(err)
            }
            req, err := http.NewRequest("PUT", "/api/transaction/input", bytes.NewBuffer(payload))
            if err != nil {
                t.Fatal(err)
            }

            rr := httptest.NewRecorder()
            router.InputHandler(rr, req)
            assert.Equal(t, tc.expectedStatus, rr.Code, "Unexpected HTTP status code")
        })
    }
}

func TestOutputHandler_ServiceErrors(t *testing.T) {
    testCases := []struct {
        name           string
        errorToReturn  error
        expectedStatus int
    }{
        {
            name:           "Service returns underfunded error",
            errorToReturn:  fmt.Errorf(utils.LogUnderfundedError),
            expectedStatus: http.StatusUnprocessableEntity,
        },
        {
            name:           "Service returns user not found error",
            errorToReturn:  fmt.Errorf(authUtils.LogUserNotFoundError),
            expectedStatus: http.StatusUnauthorized,
        },
        {
            name:           "Service returns some other error",
            errorToReturn:  fmt.Errorf("Some"),
            expectedStatus: http.StatusInternalServerError,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            mockCtrl := gomock.NewController(t)
            defer mockCtrl.Finish()

            mockService := transactioninterfaces.NewMockITransactionService(mockCtrl)
            router := &Delivery{
                service: mockService,
            }

            testTransaction := utils.OutputTransaction{UserID: "73a94f3a-42cb-11ee-9b4a-0242c0a81003", Amount: 10, Link: "http://example.com"}
            var dummyStatus utils.StatusTransaction
            mockService.EXPECT().OutputMoney(testTransaction.UserID, testTransaction.Amount, testTransaction.Link).Return(dummyStatus, tc.errorToReturn)

            payload, err := easyjson.Marshal(testTransaction)
            if err != nil {
                t.Fatal(err)
            }

            req, err := http.NewRequest("PUT", "/api/transaction/output", bytes.NewBuffer(payload))
            if err != nil {
                t.Fatal(err)
            }

            rr := httptest.NewRecorder()
            router.OutputHandler(rr, req)
            assert.Equal(t, tc.expectedStatus, rr.Code, "Unexpected HTTP status code")
        })
    }
}

