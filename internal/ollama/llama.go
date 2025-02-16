package ollama

const (
	DefaultLocalBaseURL = "http://localhost:11434"

	generateEndPoint = "/api/generate"
)

type Client interface {
	Generate(prompt string) (respbody, error)
}

func NewDefaultClient(duration int, baseURL string, model string) Client {
	return &defaultClient{
		baseURL: baseURL,
		model:   model,
		timeout: duration,
	}
}
