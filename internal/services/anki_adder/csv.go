package anki_adder

type AnkiAdderFromCSV struct {
	ankiAdderClient  ankiAdderClient
	csvNoteExtracter csvNoteExtracter
	filePath         string
}

func NewAnkiAdderFromCSV(
	ankiAdderClient ankiAdderClient,
	csvNoteExtracter csvNoteExtracter,
	csvString string,
) *AnkiAdderFromCSV {
	return &AnkiAdderFromCSV{
		ankiAdderClient: ankiAdderClient,
		csvNoteExtracter: csvNoteExtracter,
		filePath:        csvString,
	}
}

func (s *AnkiAdderFromCSV) AddNotes() error {
	notes, err := s.csvNoteExtracter.GetNotesFromCSVByFilePath(s.filePath)
	if err != nil {
		return err
	}

	return s.ankiAdderClient.AddNotes(notes)
}
