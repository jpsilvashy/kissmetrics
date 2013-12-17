package kissmetrics

import (
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	ApiKey string
}

// TODO Setup event
type Event struct {
	Person string // _p param
	Name   string // _n param
}

//
func NewClient(ApiKey string) *Client {
	return &Client{ApiKey}
}

// TODO: RecordEvent(e Event) (resp type?, err Error)
func (c *Client) RecordEvent(p, n string) (resp *http.Response, err error) {
	base := "http://trk.kissmetrics.com/e"

	// TODO this should be an event struct
	v := url.Values{}
	v.Set("_k", c.ApiKey)
	v.Set("_p", p)
	v.Set("_n", n)

	// Encode url
	queryString := v.Encode()
	request := base + "?" + queryString

	fmt.Println(request) // lazy way to log FIXME
	return http.Get(request)
}
