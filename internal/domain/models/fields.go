package models

import "strings"

type Fields map[string]string

func (f Fields) Map() map[string]string {
	return f
}

func (f Fields) ConvertHTMLTags() Fields {
	if f == nil {
		return f
	}

	m := f.Map()
	for key, val := range m {
		m[key] = strings.ReplaceAll(val, "\n", "<br>")
	}

	return m
}
