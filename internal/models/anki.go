package models

import "github.com/fatih/structs"

type Word struct {
	ModelName string
	Data      WordData
}

func (w Word) Sentence() string {
	return w.Data.getSentence()
}

func (w Word) Map() map[string]string {
	return w.Data.toMap()
}

type WordData struct {
	Sentence *string
	Target   *string
	Meaning  *string
	IPA      *string
}

func (wd WordData) getSentence() string {
	return *wd.Sentence
}

func (wd WordData) toMap() map[string]string {
	row := structs.Map(wd)

	result := make(map[string]string)
	for k, v := range row {
		if v == nil {
			continue
		}
		str, ok := v.(*string)
		if !ok {
			return nil
		}
		if str == nil {
			continue
		}
		result[k] = *str
	}

	return result
}
