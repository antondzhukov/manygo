package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

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
