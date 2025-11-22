package models

type Fields map[string]string

func (f Fields) Map() map[string]string {
	return f
}
