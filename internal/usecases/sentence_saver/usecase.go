package sentence_saver

import (
	"regexp"
	"strings"
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

// ClearString clear a string
// "some sentence" -> "some sentence"
// " some sentence    " -> "some sentence"
// "\n\nsome sentence\n\n" -> "some sentence"
// "some\n  \n sentence" -> "some sentence"
// "some      sentence" -> "some sentence"
func ClearString(s string) string {
	s = strings.TrimSpace(s)

	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\\n", " ")

	var multiSpaceRe = regexp.MustCompile(` +`)
	s = multiSpaceRe.ReplaceAllString(s, " ")

	return s
}
