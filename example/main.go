package main

import (
	"github.com/iron-io/golog"
	"github.com/jpsilvashy/kissmetrics"
	"time"
)

// FIXME, apikey should be env var
func main() {
	c := kissmetrics.NewClient("apikey")

	// Send 5 events
	for i := 0; i < 5; i++ {
		go func() {
			resp, err := c.RecordEvent("person@gmail.com", "test")
			if err != nil {
				golog.Errorln(err)
			}
			golog.Infoln(resp.StatusCode)
		}()
	}

	time.Sleep(1 * time.Second)
}
