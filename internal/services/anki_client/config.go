package anki_client

type Config struct {
	Deck      string
	NoteModel string
}

func (c Config) IsValid() bool {
	return c.Deck != "" && c.NoteModel != ""
}
