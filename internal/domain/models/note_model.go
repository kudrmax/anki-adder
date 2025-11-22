package models

import "errors"

type NoteModel string

func NewNoteModel(value string) NoteModel {
	return NoteModel(value)
}

func (nm NoteModel) IsValid() error {
	if nm == "" {
		return errors.New("note model cannot be empty")
	}

	return nil
}

func (nm NoteModel) String() string {
	return string(nm)
}
