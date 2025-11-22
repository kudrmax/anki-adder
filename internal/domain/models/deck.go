package models

import "errors"

type Deck string

func NewDeck(value string) Deck {
	return Deck(value)
}

func (d Deck) IsValid() error {
	if d == "" {
		return errors.New("deck cannot be empty")
	}

	return nil
}

func (d Deck) String() string {
	return string(d)
}
