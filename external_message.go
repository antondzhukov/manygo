package manygo

type ExternalMessageBuilder struct{}

type ExternalMessage struct {
	Url     string      `json:"url"`
	Method  string      `json:"method"`
	Headers interface{} `json:"headers,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	Timeout int64       `json:"timeout,omitempty"`
}

func (builder *ExternalMessageBuilder) BuildExternalMessage(url string, method string, headers interface{}, payload interface{}, timeout int64) *ExternalMessage {
	externalMessage := &ExternalMessage{
		Url:    url,
		Method: method,
	}

	if nil != headers {
		externalMessage.Headers = headers
	}

	if nil != payload {
		externalMessage.Payload = payload
	}

	if timeout > 0 {
		externalMessage.Timeout = timeout
	}

	return externalMessage
}
