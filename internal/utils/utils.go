package utils

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

func HttpPost(url string, body interface{}, headers map[string]string) (res *resty.Response, err error) {
	client := resty.New()

	bodyStr, err := json.Marshal(body)
	if err != nil {
		return
	}

	req := client.R()
	if len(headers) == 0 {
		req.SetHeader("Content-Type", "application/json")
	} else {
		for key, value := range headers {
			req.SetHeader(key, value)
		}
	}

	req.SetBody(bodyStr)

	return req.Post(url)
}
