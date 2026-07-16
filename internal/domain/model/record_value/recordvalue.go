package domain_model_record_value

import (
	"errors"
	"strings"

	verificationHash "github.com/Westar35/dns-check-service/internal/domain/model/verification_hash"
)

const recordValuePrefix = "dns-check-verification="

type RecordValue struct {
	value string
}

func NewRecordValueFor(hash verificationHash.VerificationHash) (RecordValue, error) {
	return NewRecordValue(recordValuePrefix + hash.String())
}

func NewRecordValue(raw string) (RecordValue, error) {
	r := RecordValue{value: raw}
	if err := r.validate(); err != nil {
		return RecordValue{}, err
	}
	return r, nil
}

func (m RecordValue) validate() error {
	if !strings.HasPrefix(m.value, recordValuePrefix) {
		return errors.New("record value: invalid prefix")
	}
	hash := strings.TrimPrefix(m.value, recordValuePrefix)
	if len(hash) != 64 {
		return errors.New("record value: invalid length")
	}
	return nil
}

func (m RecordValue) String() string { return m.value }
