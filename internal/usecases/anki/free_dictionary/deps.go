package free_dictionary

import "my/addToAnki/internal/domain/models"

type dataGetter interface {
	Get(word string) (*models.FreeDictionaryData, error)
}
