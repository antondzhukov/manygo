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
