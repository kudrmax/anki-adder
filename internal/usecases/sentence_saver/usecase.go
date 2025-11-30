package sentence_saver

import (
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

type UseCase struct {
	repo SentenceRepository
}

func New(repo SentenceRepository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (uc *UseCase) Save(sentence string) error {
	sentence = ClearString(sentence)

	if sentence == "" {
		return nil
	}

	return uc.repo.Save(sentence)
}

// Get возвращает первые n строк
// Если количество строк равно m, где m < n, то вернет m строк
func (uc *UseCase) Get(n int) ([]string, error) {
	sentences, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return sentences[:min(len(sentences), n)], nil
}

// Copy копирует первые n строк
// Если количество строк равно m, где m < n, то скопирует m строк
func (uc *UseCase) Copy(n int) error {
	sentences, err := uc.Get(n)
	if err != nil {
		return err
	}

	res := strings.Join(sentences, "\n\n")

	return clipboard.WriteAll(res)
}

func (uc *UseCase) DeleteFirstNLines(n int) error {
	return uc.repo.DeleteFirstNLines(n)
}

// ClearString clear a string
// "some sentence" -> "some sentence"
// " some sentence    " -> "some sentence"
// "\n\nsome sentence\n\n" -> "some sentence"
// "some\n  \n sentence" -> "some sentence"
// "some      sentence" -> "some sentence"
// TODO: как будто бы этой функции тут не место, нврн нужен какой-то StringCleaner
func ClearString(s string) string {
	s = strings.TrimSpace(s)

	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\\n", " ")

	var multiSpaceRe = regexp.MustCompile(` +`)
	s = multiSpaceRe.ReplaceAllString(s, " ")

	return s
}
