package kissmetrics

import (
    "testing"
)

func TestExample(t *testing.T) {
    c := Client{"example project", "example apikey"}
    sum := c.example(1, 2, 3)
    if sum != 6 {
        t.Error("expected 6 but got: ", sum)
    }
}
