package manygo

const (
	QuickReplyTypeNode                 string = "node"
	QuickReplyTypeFlow                 string = "flow"
	QuickReplyTypeDynamicBlockCallback string = "dynamic_block_callback"
)

type QuickReplyBuilder struct{}

type QuickReply struct {
	Type    string      `json:"type"`
	Caption string      `json:"caption"`
	Target  string      `json:"target,omitempty"`
	Url     string      `json:"url,omitempty"`
	Method  string      `json:"method,omitempty"`
	Headers interface{} `json:"headers,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func (builder *QuickReplyBuilder) BuildNodeQuickReply(caption string, target string) *QuickReply {
	return &QuickReply{
		Type:    QuickReplyTypeNode,
		Caption: caption,
		Target:  target,
	}
}

func (builder *QuickReplyBuilder) BuildFlowQuickReply(caption string, target string) *QuickReply {
	return &QuickReply{
		Type:    QuickReplyTypeFlow,
		Caption: caption,
		Target:  target,
	}
}

func (builder *QuickReplyBuilder) BuildDynamicBlockCallbackQuickReply(caption string, url string, method string, headers interface{}, payload interface{}) *QuickReply {
	quickReply := &QuickReply{
		Type:    QuickReplyTypeDynamicBlockCallback,
		Caption: caption,
		Url:     url,
		Method:  method,
	}

	if nil != headers {
		quickReply.Headers = headers
	}

	if nil != payload {
		quickReply.Payload = payload
	}

	return quickReply
}
