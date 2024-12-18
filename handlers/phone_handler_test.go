package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPhoneNumberHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Define test cases
	tests := []struct {
		phoneNumber  string
		countryCode  string
		expectedCode int
		expectedBody string
	}{
		{phoneNumber: "+34915872200", countryCode: "", expectedCode: http.StatusOK, expectedBody: `"phoneNumber":"+34915872200"`},
		{phoneNumber: "915872200", countryCode: "ES", expectedCode: http.StatusOK, expectedBody: `"phoneNumber":"+34915872200"`},
		{phoneNumber: "invalidNumber", countryCode: "", expectedCode: http.StatusBadRequest, expectedBody: `"message":"invalid space or character usage in phone number"`},
		{phoneNumber: "34915872200", countryCode: "", expectedCode: http.StatusOK, expectedBody: `"phoneNumber":"+34915872200"`},
		{phoneNumber: "+34 915 872 200", countryCode: "", expectedCode: http.StatusBadRequest, expectedBody: `"message":"invalid space or character usage in phone number"`},
		{phoneNumber: "2125690123", countryCode: "USS", expectedCode: http.StatusBadRequest, expectedBody: `"message":"invalid country code format, must be ISO 3166-1 alpha-2"`},
		{phoneNumber: "63131128150", countryCode: "", expectedCode: http.StatusBadRequest, expectedBody: `"message":"invalid phone number, required value is missing"`},
	}

	for _, test := range tests {
		// Create a test HTTP request
		req, _ := http.NewRequest(http.MethodGet, "/v1/phone-numbers", nil)
		q := req.URL.Query()
		q.Add("phoneNumber", test.phoneNumber)
		q.Add("countryCode", test.countryCode)
		req.URL.RawQuery = q.Encode()

		// Record the response
		w := httptest.NewRecorder()

		// Create a gin context with the request
		_, r := gin.CreateTestContext(w)
		r.GET("/v1/phone-numbers", PhoneNumberHandler)

		// Perform the request
		r.ServeHTTP(w, req)

		// Check status code
		assert.Equal(t, test.expectedCode, w.Code)

		// Check the response body
		assert.Contains(t, w.Body.String(), test.expectedBody)
	}
}
