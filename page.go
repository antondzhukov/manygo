package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

type PageService struct {
	sling        *sling.Sling
	Tags         *TagService
	GrowthTools  *GrowthToolsService
	Flows        *FlowService
	OtnTopics    *OtnTopicService
	CustomFields *CustomFieldService
	BotFields    *BotFieldService
}

func NewPageService(sling *sling.Sling) *PageService {
	base := sling.Path("page/")

	return &PageService{
		sling:        base,
		Tags:         newTagService(base),
		GrowthTools:  newGrowthToolsService(base),
		Flows:        newFlowService(base),
		OtnTopics:    newOtnTopicServiceService(base),
		CustomFields: newCustomFieldService(base),
		BotFields:    newBotFieldService(sling),
	}
}

type Page struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	AvatarLink  string `json:"avatar_link"`
	Username    string `json:"username"`
	About       string `json:"about"`
	Description string `json:"description"`
	IsPro       bool   `json:"is_pro"`
	Timezone    string `json:"timezone"`
}

type getInfoResponse struct {
	Status string `json:"status"`
	Data   Page   `json:"data,omitempty"`
}

func (s *PageService) GetInfo() (*Page, *http.Response, error) {
	getInfoResponse := new(getInfoResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("getInfo").Receive(getInfoResponse, apiError)

	return &getInfoResponse.Data, resp, MatchFirstError(err, apiError)
}

type TagService struct {
	sling *sling.Sling
}

func newTagService(sling *sling.Sling) *TagService {
	return &TagService{
		sling: sling,
	}
}

type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type getTagsResponse struct {
	Status string `json:"status"`
	Data   []Tag  `json:"data,omitempty"`
}

func (s *TagService) GetTags() (*[]Tag, *http.Response, error) {
	apiResponse := new(getTagsResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("getTags").Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type CreateTagRequest struct {
	Name string `json:"name,omitempty"`
}

type createTagResponse struct {
	Status string `json:"status"`
	Data   struct {
		Tag Tag `json:"tag"`
	} `json:"data,omitempty"`
}

func (s *TagService) CreateTag(request CreateTagRequest) (*Tag, *http.Response, error) {
	apiResponse := new(createTagResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Post("createTag").BodyJSON(request).Receive(apiResponse, apiError)

	return &apiResponse.Data.Tag, resp, MatchFirstError(err, apiError)
}

type RemoveTagRequest struct {
	Id int64 `json:"tag_id,omitempty"`
}

func (s *TagService) RemoveTag(request RemoveTagRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "removeTag", request)
}

type RemoveTagByNameRequest struct {
	Name string `json:"tag_name,omitempty"`
}

func (s *TagService) RemoveTagByName(request RemoveTagByNameRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "removeTagByName", request)
}

type OtnTopicService struct {
	sling *sling.Sling
}

func newOtnTopicServiceService(sling *sling.Sling) *OtnTopicService {
	return &OtnTopicService{
		sling: sling,
	}
}

type OtnTopic struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type getOtnTopicsResponse struct {
	Status string     `json:"status"`
	Data   []OtnTopic `json:"data,omitempty"`
}

func (s *OtnTopicService) GetOtnTopics() (*[]OtnTopic, *http.Response, error) {
	apiResponse := new(getOtnTopicsResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("getOtnTopics").Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type GrowthToolsService struct {
	sling *sling.Sling
}

func newGrowthToolsService(sling *sling.Sling) *GrowthToolsService {
	return &GrowthToolsService{
		sling: sling,
	}
}

type GrowthTool struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type getGrowthToolsResponse struct {
	Status string       `json:"status"`
	Data   []GrowthTool `json:"data,omitempty"`
}

func (s *GrowthToolsService) GetGrowthTools() (*[]GrowthTool, *http.Response, error) {
	apiResponse := new(getGrowthToolsResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("getGrowthTools").Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

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

type FlowService struct {
	sling *sling.Sling
}

func newFlowService(sling *sling.Sling) *FlowService {
	return &FlowService{
		sling: sling,
	}
}

type Folder struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent_id"`
}

type Flow struct {
	NS       string `json:"ns"`
	Name     string `json:"name"`
	FolderId int64  `json:"folder_id"`
}

type Flows struct {
	Flows   []Flow   `json:"flows"`
	Folders []Folder `json:"folders"`
}

type getFlowsResponse struct {
	Status string `json:"status"`
	Data   Flows  `json:"data,omitempty"`
}

func (s *FlowService) GetFlows() (*Flows, *http.Response, error) {
	apiResponse := new(getFlowsResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("getFlows").Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}
