package utils

import "github.com/go-resty/resty/v2"

func HttpPost(url string, body interface{}, headers map[string]string) (res *resty.Response, err error) {
	client := resty.New()

	req := client.R()
	if len(headers) == 0 {
		req.SetHeader("Content-Type", "application/json")
	} else {
		for key, value := range headers {
			req.SetHeader(key, value)
		}
	}

	req.SetBody(body)
	return req.Post(url)
}

func HttpGet(url string, headers map[string]string) (res *resty.Response, err error) {
	client := resty.New()

	req := client.R()
	if len(headers) == 0 {
		req.SetHeader("Content-Type", "application/json")
	} else {
		for key, value := range headers {
			req.SetHeader(key, value)
		}
	}

	return req.Get(url)
}
