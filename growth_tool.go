package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

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
