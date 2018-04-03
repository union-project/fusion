package propublica

import (
	"fmt"
	"net/http"
)

func (c *Client) get(reqPath string) (*http.Response, error) {
	p := apiEndpoint + reqPath
	fmt.Println(p)
	req, err := http.NewRequest("GET", p, nil)
	if err != nil {
		return nil, err
	}

	// set api key
	req.Header.Set("X-API-KEY", c.config.APIKey)

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	return resp, err
}
