package ankiconnect

import (
	"encoding/json"

	"github.com/privatesquare/bkst-go-utils/utils/errors"
)

func NewClientError(respError *errors.RestErr) error {
	if respError == nil {
		return nil
	}
	errJson, _ := json.Marshal(respError)
	return errors.New(string(errJson))
}
