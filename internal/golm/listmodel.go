package golm

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type ListModelsDataResponse struct {
	ID      string `json:"id"`
	Object  string `json:"model"`
	OwnedBy string `json:"owned_by"`
}

type ListModelResponse struct {
	Data   []ListModelsDataResponse `json:"data"`
	Object string
}

func (c client) ListModels() (ListModelResponse, error) {
	endPoint, err := url.JoinPath(c.baseURL, "/v1/models")
	if err != nil {
		return ListModelResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet, endPoint, nil)
	if err != nil {
		return ListModelResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{
		Timeout: time.Duration(c.timeout) * time.Second,
	}
	response, err := httpClient.Do(req)
	if err != nil {
		return ListModelResponse{}, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	var listResponseModel ListModelResponse
	if err := decoder.Decode(&listResponseModel); err != nil {
		return ListModelResponse{}, err
	}

	return listResponseModel, nil
}
