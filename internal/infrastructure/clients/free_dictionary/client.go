package free_dictionary

import (
	"encoding/json"
	"fmt"
	"net/http"

	"my/addToAnki/internal/domain/models"
)

const (
	path = "https://api.dictionaryapi.dev/api/v2/entries/en/%s"
)

type Client struct {
	path string
}

func New() *Client {
	return &Client{
		path: path,
	}
}

func (c *Client) pathWithWord(word string) string {
	return fmt.Sprintf(c.path, word)
}

func (c *Client) Get(word string) (*models.FreeDictionaryData, error) {
	resp, err := http.Get(c.pathWithWord(word))
	if err != nil {
		return nil, fmt.Errorf("error on do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %w", err)
	}

	var body models.FreeDictionaryData
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("error on decode body: %w", err)
	}

	return &body, nil
}
