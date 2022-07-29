package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

type BotFieldService struct {
	sling *sling.Sling
}

func newBotFieldService(sling *sling.Sling) *BotFieldService {
	return &BotFieldService{
		sling: sling,
	}
}

type BotFieldType DataType

const (
	BotFieldTypeUndefined DataType = DataTypeUndefined
	BotFieldTypeText      DataType = DataTypeText
	BotFieldTypeNumber    DataType = DataTypeNumber
	BotFieldTypeDate      DataType = DataTypeDate
	BotFieldTypeDateTime  DataType = DataTypeDateTime
	BotFieldTypeBoolean   DataType = DataTypeBoolean
)

type BotField struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	Type        BotFieldType `json:"type"`
	Description string       `json:"description"`
	Value       interface{}  `json:"value"`
}

type getBotFieldsResponse struct {
	Status string     `json:"status"`
	Data   []BotField `json:"data,omitempty"`
}

func (s *BotFieldService) GetBotFields() (*[]BotField, *http.Response, error) {
	apiResponse := new(getBotFieldsResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("getBotFields").Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type CreateBotFieldRequest struct {
	Name        string      `json:"name"`
	Type        DataType    `json:"type"`
	Description string      `json:"description"`
	Value       interface{} `json:"value"`
}

type createBotFieldResponse struct {
	Status string `json:"status"`
	Data   struct {
		BotField BotField `json:"field"`
	} `json:"data,omitempty"`
}

func (s *BotFieldService) CreateBotField(request CreateBotFieldRequest) (*BotField, *http.Response, error) {
	apiResponse := new(createBotFieldResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Post("createBotField").BodyJSON(request).Receive(apiResponse, apiError)

	return &apiResponse.Data.BotField, resp, MatchFirstError(err, apiError)
}

type SetBotFieldRequest struct {
	FieldId    int64       `json:"field_id"`
	FieldValue interface{} `json:"field_value"`
}

func (s *BotFieldService) SetBotField(request SetBotFieldRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "setBotField", request)
}

type SetBotFieldByNameRequest struct {
	FieldName  string      `json:"field_name"`
	FieldValue interface{} `json:"field_value"`
}

func (s *BotFieldService) SetBotFieldByName(request SetBotFieldByNameRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "setBotFieldByName", request)
}

type SetBotFieldsItem struct {
	Id    int64       `json:"field_id,omitempty"`
	Name  string      `json:"field_name,omitempty"`
	Value interface{} `json:"field_value"`
}

type SetBotFieldsRequest struct {
	Items []SetBotFieldsItem `json:"fields"`
}

func (s *BotFieldService) SetBotFields(request SetBotFieldsRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "setBotFields", request)
}
