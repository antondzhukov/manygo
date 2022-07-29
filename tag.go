package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

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
