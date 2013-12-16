package main

import (
	"github.com/jpsilvashy/kissmetrics"
)

func main() {
	c := kissmetrics.NewClient("apikey")
	for i := 0; i < 5; i++ {
		go c.RecordEvent("person@gmail.com", "test")
	}
}
