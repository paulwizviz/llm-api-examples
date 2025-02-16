package ollama

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type reqbody struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type respbody struct {
	Model              string `json:"model"`
	Created            string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	Context            []int  `json:"context"`
	TotalDuration      int    `json:"total_duration"`
	LoadDuration       int    `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int    `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int    `json:"eval_duration"`
}

type defaultClient struct {
	baseURL string
	model   string
	timeout int // seconds
}

func (c defaultClient) Generate(prompt string) (respbody, error) {

	fullURL, err := url.JoinPath(c.baseURL, generateEndPoint)
	if err != nil {
		return respbody{}, err
	}

	reqbody := reqbody{
		Model:  c.model,
		Prompt: prompt,
		Stream: false,
	}

	rbody, err := json.Marshal(reqbody)
	if err != nil {
		return respbody{}, err
	}
	reqBody := bytes.NewReader(rbody)

	req, err := http.NewRequest(http.MethodPost, fullURL, reqBody)
	if err != nil {
		return respbody{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{
		Timeout: time.Duration(c.timeout) * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return respbody{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return respbody{}, err
	}

	var rpBody respbody
	err = json.Unmarshal(respBody, &rpBody)
	if err != nil {
		return respbody{}, err
	}

	return rpBody, nil
}
