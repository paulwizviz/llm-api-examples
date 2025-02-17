package ollama

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type defaultClient struct {
	baseURL string
	timeout int // seconds
}

func (c defaultClient) Generate(request RequestGenerate) (*ResponseGenerate, error) {

	endPoint, err := url.JoinPath(c.baseURL, "/api/generate")
	if err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rBody := bytes.NewReader(reqBody)

	req, err := http.NewRequest(http.MethodPost, endPoint, rBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{
		Timeout: time.Duration(c.timeout) * time.Second,
	}
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	var respBody ResponseGenerate
	err = decoder.Decode(&respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}
