package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const whatsappAPIBaseURL = "https://graph.facebook.com/v19.0"

type Client struct {
	mid        string
	token      string
	httpClient *http.Client
}

type ClientOption func(c *Client)

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

func NewClient(mid, token string, opts ...ClientOption) *Client {

	c := &Client{
		mid:        mid,
		token:      token,
		httpClient: http.DefaultClient,
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

func (client *Client) SendMessage(message *Message) (*MessageResponse, error) {

	req, err := client.newRequest(http.MethodPost, "messages", message)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %v", err)
	}

	return client.send(req)
}

func (client *Client) send(request *http.Request) (*MessageResponse, error) {
	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("http error: %v", err)
	}

	defer resp.Body.Close()

	var msgResponse MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&msgResponse); err != nil {
		return nil, err
	}

	if msgResponse.Error != nil {
		return nil, msgResponse.Error
	}

	return &msgResponse, nil
}

func (client *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {

	var bodyBuf *bytes.Buffer
	if body != nil {
		bodyBuf = new(bytes.Buffer)
		if err := json.NewEncoder(bodyBuf).Encode(body); err != nil {
			return nil, err
		}
	}

	apiURL := fmt.Sprintf("%s/%s/%s", whatsappAPIBaseURL, client.mid, path)
	request, err := http.NewRequest(method, apiURL, bodyBuf)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.token))
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}
