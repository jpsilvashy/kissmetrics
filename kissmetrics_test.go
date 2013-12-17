package kissmetrics

import (
	"bytes"
	"fmt"
	"testing"
)

var c *Client

// Set the apikey
// FIXME: make an env var?
func init() {
	c = &Client{"apikey"}
}

func TestReq(t *testing.T) {
	resp, err := c.RecordEvent("jpsilvashy@gmail.com", "test")

	// return any error
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	// assert a 200
	if resp.StatusCode != 200 {
		t.Error("unexpected error: wanted status code 200, got", err)
	}

	// read from buffer
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String() // Does a complete
	fmt.Println(s)    // not good
}
