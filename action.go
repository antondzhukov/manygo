package manygo

const (
	ActionAddTag          string = "add_tag"
	ActionRemoveTag       string = "remove_tag"
	ActionSetFieldValue   string = "set_field_value"
	ActionUnsetFieldValue string = "unset_field_value"
)

type ActionBuilder struct{}

type Action struct {
	Action    string      `json:"action"`
	TagName   string      `json:"tag_name,omitempty"`
	FieldName string      `json:"field_name,omitempty"`
	Value     interface{} `json:"value,omitempty"`
}

func (builder *ActionBuilder) BuildAddTagAction(tagName string) *Action {
	return &Action{
		Action:  ActionAddTag,
		TagName: tagName,
	}
}

func (builder *ActionBuilder) BuildRemoveTagAction(tagName string) *Action {
	return &Action{
		Action:  ActionRemoveTag,
		TagName: tagName,
	}
}

func (builder *ActionBuilder) BuildSetFieldValueAction(fieldName string, value interface{}) *Action {
	return &Action{
		Action:    ActionSetFieldValue,
		FieldName: fieldName,
		Value:     value,
	}
}

func (builder *ActionBuilder) BuildUnsetFieldValueAction(fieldName string) *Action {
	return &Action{
		Action:    ActionUnsetFieldValue,
		FieldName: fieldName,
	}
}
