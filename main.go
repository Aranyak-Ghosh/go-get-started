package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type httpClient struct {
	client *http.Client
}

type HttpClient interface {
	Get(string, map[string]string, *interface{}) error
	Post(string, map[string]string, interface{}, *interface{}) error
	Put(string, map[string]string, interface{}, *interface{}) error
	Patch(string, map[string]string, interface{}, *interface{}) error
	Del(string, map[string]string, *interface{}) error
}

func (c *httpClient) Get(endpoint string, headers map[string]string, res *interface{}) error {
	req, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return err
	}

	c.constructHeaders(req, headers)

	response, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, res)

	if err != nil {
		return err
	}

	return err
}

func (c *httpClient) Post(endpoint string, headers map[string]string, body interface{}, res *interface{}) error {
	strBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(strBody))

	if err != nil {
		return err
	}

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json"
	}

	c.constructHeaders(req, headers)

	response, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, res)

	if err != nil {
		return err
	}

	return nil
}

func (c *httpClient) Put(endpoint string, headers map[string]string, body interface{}, res *interface{}) error {
	strBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(strBody))

	if err != nil {
		return err
	}

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json"
	}

	c.constructHeaders(req, headers)

	response, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, res)

	if err != nil {
		return err
	}

	return nil
}

func (c *httpClient) Patch(endpoint string, headers map[string]string, body interface{}, res *interface{}) error {
	strBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(strBody))

	if err != nil {
		return err
	}

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json"
	}

	c.constructHeaders(req, headers)

	response, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, res)

	if err != nil {
		return err
	}

	return nil
}

func (c *httpClient) Del(endpoint string, headers map[string]string, res *interface{}) error {

	req, err := http.NewRequest("DELETE", endpoint, nil)

	if err != nil {
		return err
	}

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json"
	}

	c.constructHeaders(req, headers)

	response, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, res)

	if err != nil {
		return err
	}

	return nil
}

func (c *httpClient) constructHeaders(req *http.Request, headers map[string]string) {
	for key, val := range headers {
		req.Header.Add(key, val)
	}
}

func NewHttpClient(timeout int) HttpClient {
	client := &httpClient{}

	client.client = &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	return client
}

func main() {
	client := NewHttpClient(50)
	var data interface{}
	err := client.Get("http://127.0.0.1:8080/api/v1/", nil, &data)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", data)
}

//192.168.12.24
