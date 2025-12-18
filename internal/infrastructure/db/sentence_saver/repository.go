package sentence_saver

import (
	"os"
	"strings"
)

type Repository struct {
	filePath string
}

func New(filePath string) *Repository {
	return &Repository{
		filePath: filePath,
	}
}

func (r *Repository) Save(sentence string, target *string) error {
	f, err := os.OpenFile(
		r.filePath,
		os.O_WRONLY|os.O_APPEND,
		0o644,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	targetString := sentence
	if target != nil {
		targetString = *target
	}

	_, err = f.WriteString(sentence + ";" + targetString + "\n")
	return err
}

func (r *Repository) GetAll() ([]string, error) {
	dataBytes, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	data := string(dataBytes)
	res := strings.Split(data, "\n")
	if len(res) == 0 {
		return nil, nil
	}

	return res, nil
}

func (r *Repository) DeleteFirstNLines(n int) error {
	lines, err := r.GetAll()
	if err != nil {
		return err
	}

	result := ""
	if len(lines) > n {
		result = strings.Join(lines[n:], "\n")
	}

	return os.WriteFile(r.filePath, []byte(result), 0644)
}
