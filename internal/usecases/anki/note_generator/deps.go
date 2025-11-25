package note_generator

import "context"

type generator interface {
	Generate(ctx context.Context, prompt string, format []byte) (string, error)
}
