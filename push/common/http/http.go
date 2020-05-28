package http

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	maxTimeout      = 30
	contentTypeForm = "application/x-www-form-urlencoded"
	contentTypeJson = "application/json"
)

var clientPool = sync.Pool{New: func() interface{} {
	return newClient(maxTimeout)
}}

type Client struct {
	tracingClient *http.Client
}

func newClient(timeout int) *Client {
	if timeout <= 0 && timeout > maxTimeout {
		timeout = maxTimeout
	}
	return &Client{
		tracingClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

func NewClient(timeout int) *Client {
	obj := clientPool.Get()
	if obj == nil {
		c := newClient(timeout)
		clientPool.Put(c)
		return c
	} else {
		c := (obj).(*Client)
		c.tracingClient.Timeout = 0
		return c
	}
}

func (c *Client) Get(ctx context.Context, uri string) ([]byte, error) {
	resp, err := c.tracingClient.Get(uri)
	if err != nil {
		return nil, err
	}

	return c.readBody(resp)
}

func (c *Client) BuildRequest(ctx context.Context, method, uri string, params interface{}) (*http.Request, error) {
	var data string
	switch value := params.(type) {
	case map[string]string:
		data = c.buildData(value).Encode()
	case url.Values:
		data = value.Encode()
	case string:
		data = value
	}
	req, err := http.NewRequest(method, uri, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	return req.WithContext(ctx), nil
}

func (c *Client) Do(ctx context.Context, request *http.Request) ([]byte, error) {
	request = request.WithContext(ctx)
	resp, err := c.tracingClient.Do(request)
	if err != nil {
		return nil, err
	}
	return c.readBody(resp)
}

func (c *Client) PostForm(ctx context.Context, uri string, params interface{}) ([]byte, error) {

	return c.Post(ctx, uri, params, contentTypeForm)
}

func (c *Client) PostJson(ctx context.Context, uri string, params interface{}) ([]byte, error) {

	return c.Post(ctx, uri, params, contentTypeJson)
}

func (c *Client) buildData(params map[string]string) url.Values {
	data := url.Values{}
	for key, value := range params {
		data.Add(key, value)
	}
	return data
}

func (c *Client) readBody(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) Post(ctx context.Context, uri string, params interface{}, contextType string) ([]byte, error) {
	var data string
	switch value := params.(type) {
	case map[string]string:
		data = c.buildData(value).Encode()
	case url.Values:
		data = value.Encode()
	case string:
		data = value
	}
	resp, err := c.tracingClient.Post(uri, contextType, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	return c.readBody(resp)
}
