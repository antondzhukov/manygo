package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

type CustomFieldService struct {
	sling *sling.Sling
}

func newCustomFieldService(sling *sling.Sling) *CustomFieldService {
	return &CustomFieldService{
		sling: sling,
	}
}

const (
	CustomFieldTypeUndefined DataType = DataTypeUndefined
	CustomFieldTypeText      DataType = DataTypeText
	CustomFieldTypeNumber    DataType = DataTypeNumber
	CustomFieldTypeDate      DataType = DataTypeDate
	CustomFieldTypeDateTime  DataType = DataTypeDateTime
	CustomFieldTypeBoolean   DataType = DataTypeBoolean
)

type CustomField struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Type        DataType `json:"type"`
	Description string   `json:"description"`
}

type getCustomFieldsResponse struct {
	Status string        `json:"status"`
	Data   []CustomField `json:"data,omitempty"`
}

func (s *CustomFieldService) GetCustomFields() (*[]CustomField, *http.Response, error) {
	apiResponse := new(getCustomFieldsResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("getCustomFields").Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type СreateCustomFieldRequest struct {
	Caption     string   `json:"caption"`
	Type        DataType `json:"type"`
	Description string   `json:"description"`
}

type createCustomFieldsResponse struct {
	Status string `json:"status"`
	Data   struct {
		CustomField CustomField `json:"field"`
	} `json:"data,omitempty"`
}

func (s *CustomFieldService) CreateCustomField(request СreateCustomFieldRequest) (*CustomField, *http.Response, error) {
	apiResponse := new(createCustomFieldsResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Post("createCustomField").BodyJSON(request).Receive(apiResponse, apiError)

	return &apiResponse.Data.CustomField, resp, MatchFirstError(err, apiError)
}
