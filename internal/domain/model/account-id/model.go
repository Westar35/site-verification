package domain_model_account_id

import (
	"errors"
	"fmt"
)

var ErrInvalidModel = errors.New("invalid model")

type AccountID struct {
	value string
}

func (x AccountID) GetValue() string {
	return x.value
}

func (x AccountID) Validate() error {
	if x.value == "" {
		return ErrInvalidModel
	}

	return nil
}

func NewAccountID(raw string) (AccountID, error) {
	x := AccountID{value: raw}
	if err := x.Validate(); err != nil {
		return AccountID{}, fmt.Errorf("new account id: %w", err)
	}

	return x, nil
}
