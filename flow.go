package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

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
