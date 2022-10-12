package apierrors

import (
	"errors"
	"fmt"
	"net/http"
)

func ChainNotFoundError(chainID string) *HTTPError {
	return NewHTTPError(http.StatusNotFound, fmt.Sprintf("Chain ID: %v not found", chainID), nil)
}

func BodyIsEmptyError() *HTTPError {
	return InvalidPropertyError("body", errors.New("A valid body is required"))
}

func InvalidPropertyError(propertyName string, err error) *HTTPError {
	return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid property: %v", propertyName), err)
}

func ContractExecutionError(err error) *HTTPError {
	return NewHTTPError(http.StatusBadRequest, "Failed to execute contract request", err)
}

func InvalidOffLedgerRequestError(err error) *HTTPError {
	return NewHTTPError(http.StatusBadRequest, "Supplied offledger request is invalid", err)
}

func ReceiptError(err error) *HTTPError {
	return NewHTTPError(http.StatusBadRequest, "Failed to get receipt", err)
}

func InternalServerError(err error) *HTTPError {
	return NewHTTPError(http.StatusInternalServerError, "Unknown error has occoured", err)
}
