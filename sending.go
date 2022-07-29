package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

const (
	MessageTagConfirmedEventUpdate string = "CONFIRMED_EVENT_UPDATE"
	MessageTagPostPurchaseUpdate   string = "POST_PURCHASE_UPDATE"
	MessageTagAccountUpdate        string = "ACCOUNT_UPDATE"
	MessageTagHumanAgent           string = "HUMAN_AGENT"
	MessageTagCustomerFeedback     string = "CUSTOMER_FEEDBACK"
)

type SendingService struct {
	sling *sling.Sling
}

func NewSendingService(sling *sling.Sling) *SendingService {
	return &SendingService{
		sling: sling.Path("sending/"),
	}
}

type SendingSendContentRequest struct {
	SubscriberId int64         `json:"subscriber_id"`
	Data         *DynamicBlock `json:"data"`
	MessageTag   string        `json:"message_tag"`
	OtnTopicName string        `json:"otn_topic_name,omitempty"`
}

func (s *SendingService) SendContent(request SendingSendContentRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "sendContent", request)
}

type SendingSendContentByUserRefRequest struct {
	UserRef int64        `json:"user_ref"`
	Data    DynamicBlock `json:"data"`
}

func (s *SendingService) SendContentByUserRef(request SendingSendContentByUserRefRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "sendContentByUserRef", request)
}

type SendingSendFlowRequest struct {
	SubscriberId int64  `json:"subscriber_id"`
	FlowNS       string `json:"flow_ns"`
}

func (s *SendingService) SendFlow(request SendingSendFlowRequest) (bool, *http.Response, error) {
	return DoBoolPostRequest(s.sling, "sendFlow", request)
}
