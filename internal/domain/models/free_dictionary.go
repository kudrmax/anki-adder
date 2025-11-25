package models

type (
	FreeDictionaryData []Entry

	Entry struct {
		Word      string     `json:"word"`
		Phonetic  string     `json:"phonetic"`
		Phonetics []Phonetic `json:"phonetics"`
		Meanings  []Meaning  `json:"meanings"`
		SourceURL []string   `json:"sourceUrls"`
	}

	Phonetic struct {
		Text      string `json:"text"`
		Audio     string `json:"audio"`
		SourceURL string `json:"sourceUrl,omitempty"`
	}

	Meaning struct {
		PartOfSpeech string       `json:"partOfSpeech"`
		Definitions  []Definition `json:"definitions"`
		Synonyms     []string     `json:"synonyms"`
		Antonyms     []string     `json:"antonyms"`
	}

	Definition struct {
		Definition string   `json:"definition"`
		Synonyms   []string `json:"synonyms"`
		Antonyms   []string `json:"antonyms"`
		Example    string   `json:"example,omitempty"`
	}
)

type FreeDictionaryDataClear struct {
	Word       string `json:"word,omitempty"`
	Definition string `json:"definition,omitempty"`
	Phonetic   string `json:"phonetic,omitempty"`
}

func (fdd FreeDictionaryData) GetFirstNotNullData() (*FreeDictionaryDataClear, error) {
	var word string
	for _, e := range fdd {
		if e.Word != "" {
			word = e.Word
			break
		}
	}

	var phonetic string
	for _, e := range fdd {
		if e.Phonetic != "" {
			phonetic = e.Phonetic
			break
		}
		for _, p := range e.Phonetics {
			if p.Text != "" {
				phonetic = p.Text
				break
			}
		}
	}

	var definition string
	for _, e := range fdd {
		for _, m := range e.Meanings {
			for _, d := range m.Definitions {
				if d.Definition != "" {
					definition = d.Definition
					break
				}
			}
		}
	}

	if word == "" {
		//return nil, errors.New("word is empty")
	}
	if phonetic == "" {
		//return nil, errors.New("phonetic is empty")
	}
	if definition == "" {
		//return nil, errors.New("definition is empty")
	}

	return &FreeDictionaryDataClear{
		Word:       word,
		Phonetic:   phonetic,
		Definition: definition,
	}, nil
}
