package kissmetrics

import (
	"bytes"
	"fmt"
	"testing"
)

var c *Client

func init() {
	c = &Client{"apikey"}
}

func TestReq(t *testing.T) {
	resp, err := c.RecordEvent("jpsilvashy@gmail.com", "test")
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if resp.StatusCode != 200 {
		t.Error("unexpected error: wanted status code 200, got", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String() // Does a complete
	fmt.Println(s)
}
