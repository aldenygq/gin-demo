package tools

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)


type Client struct {
	Client *http.Client
}

func NewClient(timeout int64) *Client {
	return &Client{
		Client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
	
}

func (c *Client) Get(url string, params map[string]string, header map[string]interface{}) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	for k, v := range header {
		req.Header.Set(k, v.(string))
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	req.Close = true
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func (c *Client) GetWithCookie(url string, params map[string]string, header map[string]interface{}) ([]byte, []*http.Cookie, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	
	for k, v := range header {
		req.Header.Set(k, v.(string))
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	req.Close = true
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return dat, resp.Cookies(), nil
}

func (c *Client) PostForm(uri string, params map[string]string) ([]byte, error) {
	
	form := make(url.Values)
	for k, v := range params {
		form.Add(k, v)
	}
	resp, err := c.Client.PostForm(uri, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return dat, nil
	
}

func (c *Client) PostJson(uri string, params []byte) ([]byte, error) {
	
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(params))
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (c *Client) PutJson(uri string, params []byte, header map[string]interface{}) ([]byte, error) {
	
	req, err := http.NewRequest("PUT", uri, bytes.NewBuffer(params))
	
	for k, v := range header {
		req.Header.Set(k, v.(string))
	}
	
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}


