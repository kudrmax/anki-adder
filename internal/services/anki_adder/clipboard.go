package anki_adder

type AnkiAdderFromClipboard struct {
	ankiAdderClient ankiAdderClient
}

func NewAnkiAdderFromClipboard(ankiAdderClient ankiAdderClient) *AnkiAdderFromClipboard {
	return &AnkiAdderFromClipboard{
		ankiAdderClient: ankiAdderClient,
	}
}

func (s *AnkiAdderFromClipboard) AddNotes() error {
	return nil
}
