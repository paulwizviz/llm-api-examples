package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type defaultClient struct {
	baseURL string
	timeout int // seconds
}

func (c defaultClient) GenerateAPI(request RequestGenerate) ([]ResponseGenerate, error) {

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

	var responses []ResponseGenerate
	if request.Stream {
		responses := parseGenerateStream(response.Body)
		return responses, nil
	}

	decoder := json.NewDecoder(response.Body)
	var resp ResponseGenerate
	err = decoder.Decode(&resp)
	if err != nil {
		return nil, err
	}
	responses = append(responses, resp)
	return responses, nil
}

func parseGenerateStream(body io.ReadCloser) []ResponseGenerate {

	var responses []ResponseGenerate
	reader := bufio.NewReader(body)
loop:
	for {
		ln, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break loop
			}
			continue
		}
		if ln[0] != '{' && ln[len(ln)-1] != '}' {
			continue
		}
		var response ResponseGenerate
		err = json.Unmarshal([]byte(ln), &response)
		if err != nil {
			continue
		}
		responses = append(responses, response)
	}
	return responses
}
