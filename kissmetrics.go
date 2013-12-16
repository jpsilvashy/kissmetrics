package kissmetrics

import ()

type Client struct {
    projectId string
    apikey    string
}

func (c *Client) example(numbers ...int) int {
    var sum int
    for _, num := range numbers {
        sum += num
    }
    return sum
}
