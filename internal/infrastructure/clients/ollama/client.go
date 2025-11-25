package ollama

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

type Config struct {
	model  string
	stream bool
}

type Client struct {
	client *api.Client
	config Config
}

func NewClient(model string, stream bool) (*Client, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
		config: Config{
			model:  model,
			stream: stream,
		},
	}, nil
}
func (c *Client) Generate(ctx context.Context, prompt string, format []byte) (string, error) {
	var result string

	err := c.client.Generate(ctx, &api.GenerateRequest{
		Model:  c.config.model,
		Stream: &c.config.stream,
		Prompt: prompt,
		Format: format,
	}, func(res api.GenerateResponse) error {
		result = res.Response
		fmt.Print(res.Response)
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error on generate: %w", err)
	}

	return result, nil
}
