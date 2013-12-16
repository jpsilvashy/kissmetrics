package kissmetrics

import (
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	ApiKey string
}

type Event struct {
	Person string
	N      string
}

func NewClient(ApiKey string) *Client {
	return &Client{ApiKey}
}

func (c *Client) RecordEvent(p, n string) (resp *http.Response, err error) {
	base := "http://trk.kissmetrics.com/e"
	v := url.Values{}
	v.Set("_k", c.ApiKey)
	v.Set("_p", p)
	v.Set("_n", n)
	queryString := v.Encode()
	request := base + "?" + queryString
	fmt.Println(request)
	return http.Get(request)
}
