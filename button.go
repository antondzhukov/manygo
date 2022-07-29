package manygo

const (
	ButtonTypeCall                 string = "call"
	ButtonTypeUrl                  string = "url"
	ButtonTypeFlow                 string = "flow"
	ButtonTypeNode                 string = "node"
	ButtonTypeBuy                  string = "buy"
	ButtonTypeDynamicBlockCallback string = "dynamic_block_callback"
)

const (
	WebViewSizeFull    string = "full"
	WebViewSizeMedium  string = "medium"
	WebViewSizeCompact string = "compact"
)

type ButtonsBuilder struct{}

type Button struct {
	Type          string      `json:"type"`
	Caption       string      `json:"caption"`
	Phone         string      `json:"phone,omitempty"`
	Url           string      `json:"url,omitempty"`
	Method        string      `json:"method,omitempty"`
	WebViewSize   string      `json:"web_view_size,omitempty"`
	Target        string      `json:"target,omitempty"`
	Customer      Customer    `json:"customer,omitempty"`
	Product       Product     `json:"product,omitempty"`
	SuccessTarget string      `json:"success_target,omitempty"`
	Headers       interface{} `json:"headers,omitempty"`
	Payload       interface{} `json:"payload,omitempty"`
}

type Customer struct {
	ShippingAddress bool `json:"shipping_address"`
	ContactName     bool `json:"contact_name"`
	ContactPhone    bool `json:"contact_phone"`
	ContactEmail    bool `json:"contact_email"`
}

type Product struct {
	Label string `json:"label"`
	Cost  int64  `json:"cost"`
}

func (builder *ButtonsBuilder) BuildCallButton(caption string, phone string) *Button {
	return &Button{
		Type:    ButtonTypeCall,
		Caption: caption,
		Phone:   phone,
	}
}

func (builder *ButtonsBuilder) BuildUrlButton(caption string, url string, webViewSize string) *Button {
	button := &Button{
		Type:    ButtonTypeUrl,
		Caption: caption,
		Url:     url,
	}

	if webViewSize != "" {
		button.WebViewSize = webViewSize
	}

	return button
}

func (builder *ButtonsBuilder) BuildFlowButton(caption string, target string) *Button {
	return &Button{
		Type:    ButtonTypeFlow,
		Caption: caption,
		Target:  target,
	}
}

func (builder *ButtonsBuilder) BuildNodeButton(caption string, target string) *Button {
	return &Button{
		Type:    ButtonTypeNode,
		Caption: caption,
		Target:  target,
	}
}

func (builder *ButtonsBuilder) BuildBuyButton(caption string, customer Customer, product Product, successTarget string) *Button {
	return &Button{
		Type:          ButtonTypeBuy,
		Caption:       caption,
		Customer:      customer,
		Product:       product,
		SuccessTarget: successTarget,
	}
}

func (builder *ButtonsBuilder) BuildCallbackButton(caption string, url string, method string, headers interface{}, payload interface{}) *Button {
	button := &Button{
		Type:    ButtonTypeDynamicBlockCallback,
		Caption: caption,
		Url:     url,
		Method:  method,
	}

	if nil != headers {
		button.Headers = headers
	}

	if nil != payload {
		button.Payload = payload
	}

	return button
}
