package helpers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrNoRecordFound is an error indicating that no record was found for a specific ID.
// This can be used to signal that a request for a record failed because the record doesn't exist.
var ErrNoRecordFound = errors.New("services: There is no record with this id")

// ErrUnprocessableEntity is an error indicating that the server cannot process the request due to semantic errors.
// This can be returned when the server understands the content type of the request entity but is unable to process the contained instructions.
var ErrUnprocessableEntity = errors.New(http.StatusText(http.StatusUnprocessableEntity))

// MissingFieldErr is a helper function that formats an error message when one or more required fields are missing.
// It sends a 400 Bad Request status with a message listing the missing fields.
func MissingFieldErr(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Missing required field(s): " + message, // Include the specific missing fields in the response
	})
}

// BadRequestError is a helper function to return a standard 400 Bad Request error message.
// It sends a generic "Bad Request" message in response when the server cannot process the request due to malformed syntax.
func BadRequestError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": http.StatusText(http.StatusBadRequest), // Standard HTTP Bad Request message
	})
}

// ServerError is a helper function that handles unexpected server errors. It logs the error and returns a generic error message to the client.
// It ensures that sensitive details about the error are not exposed to the client.
func ServerError(c *gin.Context, err error) {
	log.Println(err) // Log the detailed error for internal debugging purposes
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "An unexpected server error occurred. Please try again later.", // Return a generic message to the client
	})
}

// NotFoundError is a helper function to return a 404 Not Found error in JSON format.
func NotFoundError(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": message,
	})
}

// NotFoundPage sends a 404 response with a default "Page not found" message in JSON format.
func NotFoundPage(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": http.StatusText(http.StatusNotFound),
	})
}
