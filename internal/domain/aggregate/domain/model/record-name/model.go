package domain_aggregate_domain_model_record_name

import (
	"errors"
	"fmt"

	"github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/domain-name"
)

const recordNamePrefix = "_dns-check."

type RecordName struct {
	value string
}

func (x RecordName) GetValue() string {
	return x.value
}

func (x RecordName) Validate() error {
	if x.value == "" {
		return errors.New("invalid model")
	}

	return nil
}

func NewRecordName(raw string) (RecordName, error) {
	x := RecordName{value: raw}
	if err := x.Validate(); err != nil {
		return RecordName{}, fmt.Errorf("new record name: %w", err)
	}

	return x, nil
}

func NewRecordNameFor(domainName domain_aggregate_domain_model_domain_name.DomainName) (RecordName, error) {
	x, err := NewRecordName(recordNamePrefix + domainName.GetValue())
	if err != nil {
		return RecordName{}, fmt.Errorf("new record name for domain: %w", err)
	}

	return x, nil
}
