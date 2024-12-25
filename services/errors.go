package services

import (
	"errors"
	"net/http"
)

var ErrNoRecordFound = errors.New("services: There is no record with this id")
var ErrUnprocessableEntity = errors.New(http.StatusText(http.StatusUnprocessableEntity))