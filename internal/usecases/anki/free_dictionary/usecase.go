package free_dictionary

import (
	"fmt"

	"my/addToAnki/internal/domain/models"
)

type UseCase struct {
	dataGetter dataGetter
}

func New(dataGetter dataGetter) *UseCase {
	return &UseCase{
		dataGetter: dataGetter,
	}
}

func (uc *UseCase) GetData(word string) (*models.FreeDictionaryDataClear, error) {
	data, err := uc.dataGetter.Get(word)
	if err != nil {
		return nil, fmt.Errorf("error on getting data: %w", err)
	}

	return data.GetFirstNotNullData()
}
