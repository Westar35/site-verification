package domain_model_account_id

import (
	"errors"
)

type AccountID struct {
	value string
}

func (m AccountID) String() string {
	return m.value
}

func (m AccountID) validate() error {
	if m.value == "" {
		return errors.New("account id: value required")
	}
	return nil
}

func New(id string) (AccountID, error) {
	model := AccountID{value: id}

	if err := model.validate(); err != nil {
		return AccountID{}, err
	}
	return model, nil
}
