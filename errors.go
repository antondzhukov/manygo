package manygo

import (
	"errors"
	"fmt"
	"strconv"
)

func MatchFirstError(httpError error, errorResponse *ErrorResponse) error {
	if httpError != nil {
		return httpError
	}

	if errorResponse != nil {
		addon := ""
		if 0 != errorResponse.ErrorCode {
			addon += strconv.Itoa(errorResponse.ErrorCode) + " "
		}

		addon += fmt.Sprintf("%v", errorResponse.Details)

		return createError(errorResponse.Message + " " + addon)
	}

	return nil
}

func createError(message string) error {
	return errors.New(message)
}
