package domain_model_account_id

import (
	"errors"
	"fmt"
)

var errInvalidModel = errors.New("invalid model")

type AccountID struct {
	value string
}

func (x AccountID) GetValue() string {
	return x.value
}

func (x AccountID) validate() error {
	if x.value == "" {
		return errInvalidModel
	}

	return nil
}

func NewAccountID(raw string) (AccountID, error) {
	x := AccountID{value: raw}
	if err := x.validate(); err != nil {
		return AccountID{}, fmt.Errorf("new account id: %w", err)
	}

	return x, nil
}
