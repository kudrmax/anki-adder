package note_generator

import "encoding/json"

type FieldSchema struct {
	Type string `json:"type"`
}

type Schema struct {
	Type       string                 `json:"type"`
	Properties map[string]FieldSchema `json:"properties"`
	Required   []string               `json:"required"`
}

func BuildFormat(fields []string) ([]byte, error) {
	props := make(map[string]FieldSchema, len(fields))
	for _, f := range fields {
		props[f] = FieldSchema{Type: "string"}
	}

	schema := Schema{
		Type:       "object",
		Properties: props,
		Required:   fields,
	}

	data, err := json.MarshalIndent(schema, "", "\t")
	if err != nil {
		return nil, err
	}

	return data, nil
}
