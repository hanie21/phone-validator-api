package handlers

import (
	"PhoneValidatorAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhoneNumberResponse struct {
	PhoneNumber      string `json:"phoneNumber"`
	CountryCode      string `json:"countryCode"`
	AreaCode         string `json:"areaCode"`
	LocalPhoneNumber string `json:"localPhoneNumber"`
}

type ErrorResponse struct {
	PhoneNumber string `json:"phoneNumber"`
	Error       struct {
		Message string `json:"message"`
	} `json:"error"`
}

type PhoneNumberRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	CountryCode string `json:"countryCode"`
}

func PhoneNumberHandler(c *gin.Context) {
	phoneNumber := c.Query("phoneNumber")
	countryCode := c.Query("countryCode")

	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "phoneNumber parameter is required",
		})
		return
	}

	// Validate the phone number using the service logic
	num, err := services.ValidatePhoneNumber(phoneNumber, countryCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			PhoneNumber: phoneNumber,
			Error: struct {
				Message string `json:"message"`
			}{
				Message: err.Error(),
			},
		})
		return
	}

	// Extract components from the parsed phone number
	e164Format, countryCodeStr, areaCode, localPhoneNumber := services.ExtractPhoneNumberComponents(num)

	// Return the response
	c.JSON(http.StatusOK, PhoneNumberResponse{
		PhoneNumber:      e164Format,
		CountryCode:      countryCodeStr,
		AreaCode:         areaCode,
		LocalPhoneNumber: localPhoneNumber,
	})
}

func CreatePhoneNumberHandler(c *gin.Context) {
	var request PhoneNumberRequest

	// Bind JSON body to request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate the phone number using the service logic
	num, err := services.ValidatePhoneNumber(request.PhoneNumber, request.CountryCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			PhoneNumber: request.PhoneNumber,
			Error: struct {
				Message string `json:"message"`
			}{
				Message: err.Error(),
			},
		})
		return
	}

	// Extract components from the parsed phone number
	e164Format, countryCodeStr, areaCode, localPhoneNumber := services.ExtractPhoneNumberComponents(num)

	// Return the response
	c.JSON(http.StatusOK, PhoneNumberResponse{
		PhoneNumber:      e164Format,
		CountryCode:      countryCodeStr,
		AreaCode:         areaCode,
		LocalPhoneNumber: localPhoneNumber,
	})
}
