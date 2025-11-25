package note_generator

import (
	"context"
)

type UseCase struct {
	generator generator
	fields    []string
	format    []byte
}

func New(
	generator generator,
	fields []string,
) (*UseCase, error) {
	format, err := BuildFormat(fields)
	if err != nil {
		return nil, err
	}

	return &UseCase{
		generator: generator,
		fields:    fields,
		format:    format,
	}, nil
}

var (
	promptTask = "You are an assistant helping an English learner add vocabulary to Anki.\n\nInput:\n\nSentence: a full sentence in English.\n\nTarget: a single word or short phrase taken from that sentence, which the learner wants to study.\n\nYour tasks:\n\nMeaning (in English):\n\nExplain the Target in clear, simple English.\n\nUse the Sentence as context to choose the correct sense of the word.\n\nIf the word is idiomatic, phrasal, or metaphorical in this context, explain that specific meaning.\n\nDon’t translate into other languages; only explain in English.\n\nIPA (American English):\n\nGive the IPA transcription of the Target in American English.\n\nIf the Target is a multi-word expression, give the IPA for the whole expression as it would naturally be pronounced in American English.\n\nImportant guidelines:\n\nAlways base your explanation on the context of the Sentence, not just dictionary definitions.\n\nBe concise but clear and learner-friendly.\n\nIf the Target can be different parts of speech, choose the one that fits the Sentence and explain that."
)

func (uc *UseCase) GenerateNote(sentence, target string) (string, error) {
	prompt := promptTask + "\n\nInput:\n\nSentence: " + sentence + "\n\nTarget: " + target

	return uc.generator.Generate(context.Background(), prompt, uc.format)
}
