package anki_adder

type AnkiAdderFromClipboard struct {
	ankiAdderClient  ankiAdderClient
	csvNoteExtracter csvNoteExtracter
	data             string
}

func NewAnkiAdderFromClipboard(
	ankiAdderClient ankiAdderClient,
	csvNoteExtracter csvNoteExtracter,
	data string,
) *AnkiAdderFromClipboard {
	return &AnkiAdderFromClipboard{
		ankiAdderClient:  ankiAdderClient,
		csvNoteExtracter: csvNoteExtracter,
		data:             data,
	}
}

func (s *AnkiAdderFromClipboard) AddNotes() error {
	notes, err := s.csvNoteExtracter.GetNotesFromCSVByString(s.data)
	if err != nil {
		return err
	}

	return s.ankiAdderClient.AddNotes(notes)
}
