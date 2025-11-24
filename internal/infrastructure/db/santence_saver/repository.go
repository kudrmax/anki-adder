package santence_saver

import (
	"os"
)

type Repository struct {
	filePath string
}

func New(filePath string) *Repository {
	return &Repository{
		filePath: filePath,
	}
}

func (r *Repository) Save(sentence string) error {
	f, err := os.OpenFile(
		r.filePath,
		os.O_WRONLY|os.O_APPEND,
		0o644,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(sentence + "\n")
	return err
}
