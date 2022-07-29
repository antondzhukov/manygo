package manygo

import (
	"github.com/dghubble/sling"
	"net/http"
)

const manychatAPI = "https://api.manychat.com/fb/"

type Client struct {
	sling        *sling.Sling
	httpClient   *http.Client
	url          string
	bearer       string
	Page         *PageService
	Subscribers  *SubscriberService
	Sending      *SendingService
	DynamicBlock *DynamicBlockBuilder
}

func NewClient(httpClient *http.Client, bearer string) *Client {
	client := Client{
		httpClient: httpClient,
		bearer:     bearer,
		url:        manychatAPI,
	}

	client.constructClient()

	return &client
}

func (client *Client) SetApiUrl(url string) {
	client.url = url
	client.constructClient()
}

func (client *Client) SetBearer(bearer string) {
	client.bearer = bearer
	client.constructClient()
}

func (client *Client) getBase() *sling.Sling {
	base := sling.New().Client(client.httpClient).Base(client.url)
	base.Set("Authorization", "Bearer "+client.bearer)
	base.Set("accept", "application/json")

	return base
}

func (client *Client) constructClient() {
	client.sling = client.getBase()
	client.Page = NewPageService(client.sling.New())
	client.Subscribers = NewSubscriberService(client.sling.New())
	client.Sending = NewSendingService(client.sling.New())
	client.DynamicBlock = NewDynamicBlockBuilder()
}
