package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

type SubscriberService struct {
	sling *sling.Sling
}

func NewSubscriberService(sling *sling.Sling) *SubscriberService {
	return &SubscriberService{
		sling: sling.Path("subscriber/"),
	}
}

type Subscriber struct {
	Id                string                   `json:"id"`
	PageId            string                   `json:"page_id"`
	UserRefs          []SubscriberRefField     `json:"user_refs,omitempty"`
	FirstName         string                   `json:"first_name"`
	LastName          string                   `json:"last_name"`
	Name              string                   `json:"name"`
	Gender            string                   `json:"gender"`
	ProfilePic        string                   `json:"profile_pic"`
	Locale            string                   `json:"locale"`
	Language          string                   `json:"language"`
	TimeZone          string                   `json:"time_zone"`
	LiveChatUrl       string                   `json:"live_chat_url"`
	LastInputText     string                   `json:"last_input_text"`
	OptInPhone        bool                     `json:"optin_phone"`
	Phone             string                   `json:"phone"`
	OptInEmail        bool                     `json:"optin_email"`
	Email             string                   `json:"email"`
	Subscribed        string                   `json:"subscribed"`
	LastInteraction   string                   `json:"last_interaction"`
	LastSeen          string                   `json:"last_seen"`
	IsFollowupEnabled bool                     `json:"is_followup_enabled"`
	IgUsername        string                   `json:"ig_username"`
	IgId              int64                    `json:"ig_id"`
	WhatsAppPhone     string                   `json:"whatsapp_phone"`
	OptInWhatsApp     bool                     `json:"optin_whatsapp"`
	CustomFields      []SubscriberCustomField  `json:"custom_fields"`
	ShopifyFields     []SubscriberShopifyField `json:"shopify_fields,omitempty"`
	Tags              []Tag                    `json:"tags,omitempty"`
}

type SubscriberRefField struct {
	UserRef string `json:"user_ref"`
	OptedIn string `json:"opted_in"`
}

type SubscriberCustomField struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Type        DataType    `json:"type"`
	Description string      `json:"description"`
	Value       interface{} `json:"value"`
}

type SubscriberShopifyField struct {
	Id                    int64    `json:"id"`
	State                 string   `json:"state"`
	Currency              string   `json:"currency"`
	TotalSpent            string   `json:"total_spent"`
	OrdersCount           int64    `json:"orders_count"`
	AcceptsMarketing      bool     `json:"accepts_marketing"`
	LastOrderId           int64    `json:"last_order_id"`
	LastOrderCreatedAt    string   `json:"last_order_created_at"`
	Tags                  []string `json:"tags"`
	LastCheckoutId        int64    `json:"last_checkout_id"`
	LastCheckoutPrice     string   `json:"last_checkout_price"`
	LastCheckoutCreatedAt string   `json:"last_checkout_created_at"`
}

type singleSubscriberInfoResponse struct {
	Status string     `json:"status"`
	Data   Subscriber `json:"data,omitempty"`
}

type multipleSubscriberResponse struct {
	Status string       `json:"status"`
	Data   []Subscriber `json:"data,omitempty"`
}

type GetSubscriberInfoRequest struct {
	SubscriberId int64 `url:"subscriber_id"`
}

func (s *SubscriberService) GetInfo(request GetSubscriberInfoRequest) (*Subscriber, *http.Response, error) {
	apiResponse := new(singleSubscriberInfoResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("getInfo").QueryStruct(&request).Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type FindSubscriberByNameRequest struct {
	Name string `url:"name"`
}

func (s *SubscriberService) FindByName(request FindSubscriberByNameRequest) (*[]Subscriber, *http.Response, error) {
	apiResponse := new(multipleSubscriberResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("findByName").QueryStruct(&request).Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type GetSubscriberInfoByUserRefRequest struct {
	UserRef int64 `url:"user_ref"`
}

func (s *SubscriberService) GetInfoByUserRef(request GetSubscriberInfoByUserRefRequest) (*Subscriber, *http.Response, error) {
	apiResponse := new(singleSubscriberInfoResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("getInfoByUserRef").QueryStruct(&request).Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type FindSubscriberByCustomFieldRequest struct {
	FieldId    int64       `url:"field_id"`
	FieldValue interface{} `url:"field_value"`
}

func (s *SubscriberService) FindByCustomField(request FindSubscriberByCustomFieldRequest) (*[]Subscriber, *http.Response, error) {
	apiResponse := new(multipleSubscriberResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("findByCustomField").QueryStruct(&request).Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type FindSubscriberByEmailRequest struct {
	Email string `url:"email"`
}

func (s *SubscriberService) FindByEmail(request FindSubscriberByEmailRequest) (*Subscriber, *http.Response, error) {
	apiResponse := new(singleSubscriberInfoResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("findBySystemField").QueryStruct(&request).Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type FindSubscriberByPhoneRequest struct {
	Phone string `url:"phone"`
}

func (s *SubscriberService) FindByPhone(request FindSubscriberByPhoneRequest) (*Subscriber, *http.Response, error) {
	apiResponse := new(singleSubscriberInfoResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Get("findBySystemField").QueryStruct(&request).Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type AddSubscriberTagRequest struct {
	SubscriberId int64 `json:"subscriber_id"`
	TagId        int64 `json:"tag_id"`
}

func (s *SubscriberService) AddTag(request AddSubscriberTagRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "addTag", request)
}

type AddSubscriberTagByNameRequest struct {
	SubscriberId int64  `json:"subscriber_id"`
	TagName      string `json:"tag_name"`
}

func (s *SubscriberService) AddTagByName(request AddSubscriberTagRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "addTagByName", request)
}

type RemoveSubscriberTagRequest struct {
	SubscriberId int64 `json:"subscriber_id"`
	TagId        int64 `json:"tag_id"`
}

func (s *SubscriberService) RemoveTag(request RemoveSubscriberTagRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "removeTag", request)
}

type RemoveSubscriberTagByNameRequest struct {
	SubscriberId int64  `json:"subscriber_id"`
	TagName      string `json:"tag_name"`
}

func (s *SubscriberService) RemoveTagByName(request RemoveSubscriberTagByNameRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "removeTagByName", request)
}

type SetSubscriberCustomFieldRequest struct {
	SubscriberId int64       `json:"subscriber_id"`
	FieldId      int64       `json:"field_id"`
	FieldValue   interface{} `json:"field_value"`
}

func (s *SubscriberService) SetCustomField(request SetSubscriberCustomFieldRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "setCustomField", request)
}

type SetCustomFieldsItem struct {
	Id    int64       `json:"field_id,omitempty"`
	Name  string      `json:"field_name,omitempty"`
	Value interface{} `json:"field_value"`
}

type SetCustomFieldsRequest struct {
	Items []SetCustomFieldsItem `json:"fields"`
}

func (s *BotFieldService) SetCustomFields(request SetCustomFieldsRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "setCustomFields", request)
}

type SetSubscriberCustomFieldByNameRequest struct {
	SubscriberId int64       `json:"subscriber_id"`
	FieldName    string      `json:"field_name"`
	FieldValue   interface{} `json:"field_value"`
}

func (s *SubscriberService) SetCustomFieldByName(request SetSubscriberCustomFieldByNameRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "setCustomFieldByName", request)
}

type SubscriberVerifyBySignedRequestRequest struct {
	SubscriberId  int64  `json:"subscriber_id"`
	SignedRequest string `json:"signed_request"`
}

func (s *SubscriberService) VerifyBySignedRequest(request SubscriberVerifyBySignedRequestRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "verifyBySignedRequest", request)
}

type CreateSubscriberRequest struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Phone         string `json:"phone"`
	WhatsAppPhone string `json:"whatsapp_phone,omitempty"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	HasOptInSMS   bool   `json:"has_opt_in_sms"`
	HasOptInEmail bool   `json:"has_opt_in_email"`
	ConsentPhrase string `json:"consent_phrase"`
}

func (s *SubscriberService) CreateSubscriber(request CreateSubscriberRequest) (*Subscriber, *http.Response, error) {
	apiResponse := new(singleSubscriberInfoResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Post("createSubscriber").BodyJSON(request).Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}

type UpdateSubscriberRequest struct {
	SubscriberId  int64  `json:"subscriber_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Phone         string `json:"phone"`
	WhatsAppPhone string `json:"whatsapp_phone,omitempty"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	HasOptInSMS   bool   `json:"has_opt_in_sms"`
	HasOptInEmail bool   `json:"has_opt_in_email"`
	ConsentPhrase string `json:"consent_phrase"`
}

func (s *SubscriberService) UpdateSubscriber(request UpdateSubscriberRequest) (*Subscriber, *http.Response, error) {
	apiResponse := new(singleSubscriberInfoResponse)
	apiError := new(ErrorResponse)

	resp, err := s.sling.New().Post("updateSubscriber").BodyJSON(request).Receive(apiResponse, apiError)

	return &apiResponse.Data, resp, MatchFirstError(err, apiError)
}
