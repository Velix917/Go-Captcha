package gocaptcha

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const CaptchaIn string = "https://2captcha.com/in.php"
const CaptchaRes string = "https://2captcha.com/res.php"

type Client struct {
	Key     string
	HClient *http.Client
}

func NewClient(key string, hclient ...*http.Client) *Client {
	if len(hclient) > 0 {
		return &Client{Key: key, HClient: hclient[0]}
	}
	return &Client{Key: key, HClient: &http.Client{}}
}

func (r *Client) request(uri string, params string) (string, error) {
	uri = uri + "?" + params
	res, err := r.HClient.Get(uri)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (r *Client) doApiCall(endpoint string, params string) (string, error) {
	result, err := r.request(endpoint, params)
	if err != nil {
		return "", err
	}

	ok := checkApiResponse(result)
	if !ok {
		return "", fmt.Errorf("error: %s", result)
	}

	return result, nil
}

func (r *Client) getResult(reqId string) (string, error) {
	params := url.Values{}
	params.Add("id", reqId)
	params.Add("key", r.Key)
	params.Add("action", "get")

	result, err := r.doApiCall(CaptchaRes, params.Encode())
	if err != nil {
		return "", err
	}

	if result == "CAPCHA_NOT_READY" {
		return "", nil
	}

	return parseResult(result), nil
}

func (r *Client) Solve(captcha ICaptchaRequest) (string, error) {
	captcha.Add("key", r.Key)

	result, err := r.doApiCall(CaptchaIn, captcha.ToString())
	if err != nil {
		return "", err
	}

	reqId := parseResult(result)

	var counter int = 120 / 5
	for i := 0; i < counter; i++ {
		result, err = r.getResult(reqId)
		if err != nil {
			return "", err
		}

		if result != "" {
			return result, nil
		}

		time.Sleep(5 * time.Second)
	}
	return "", fmt.Errorf("Timeout")
}
