package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

func DoBoolPostRequest(sling *sling.Sling, path string, requestParams interface{}) (bool, *http.Response, error) {
	apiResponse := new(BoolSuccessResponse)
	apiError := new(ErrorResponse)

	resp, err := sling.New().Post(path).BodyJSON(requestParams).Receive(apiResponse, apiError)

	return resp != nil && resp.StatusCode == http.StatusOK, resp, MatchFirstError(err, apiError)
}
