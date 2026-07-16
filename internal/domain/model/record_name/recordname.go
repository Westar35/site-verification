package domain_model_record_name

import (
	"errors"

	domain "github.com/Westar35/dns-check-service/internal/domain/model/normalized_domain_name"
)

type RecordName struct {
	value string
}

func NewRecordNameFor(domain domain.NormalizedDomainName) (RecordName, error) {
	return NewRecordName("_dns-check." + domain.String())
}

func NewRecordName(raw string) (RecordName, error) {
	r := RecordName{value: raw}
	if err := r.validate(); err != nil {
		return RecordName{}, err
	}
	return r, nil
}

func (m RecordName) validate() error {
	if m.value == "" {
		return errors.New("record name: value required")
	}
	return nil
}

func (m RecordName) String() string { return m.value }
