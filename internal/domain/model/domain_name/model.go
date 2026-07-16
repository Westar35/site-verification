package domain_model_domain_name

import (
	"errors"
	"strings"
)

type DomainName struct {
	value string
}

func (m DomainName) String() string {
	return m.value
}

func (m DomainName) validate() error {
	if strings.TrimSpace(m.value) == "" {
		return errors.New("domain name: value required")
	}
	return nil
}

func NewDomainName(raw string) (DomainName, error) {
	d := DomainName{value: raw}
	if err := d.validate(); err != nil {
		return DomainName{}, err
	}
	return d, nil
}
