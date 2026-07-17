package domain_aggregate_domain_model_record_value

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/domain_name"
	"github.com/Westar35/site-verification/internal/domain/model/account_id"
)

const (
	recordValuePrefix = "dns-check-verification="
	hashHexLength     = 64
)

var errInvalidModel = errors.New("invalid model")

type RecordValue struct {
	value string
}

func (x RecordValue) GetValue() string {
	return x.value
}

func (x RecordValue) validate() error {
	if !strings.HasPrefix(x.value, recordValuePrefix) {
		return errInvalidModel
	}

	hash := strings.TrimPrefix(x.value, recordValuePrefix)
	if len(hash) != hashHexLength {
		return errInvalidModel
	}

	return nil
}

func NewRecordValue(raw string) (RecordValue, error) {
	x := RecordValue{value: raw}
	if err := x.validate(); err != nil {
		return RecordValue{}, fmt.Errorf("new record value: %w", err)
	}

	return x, nil
}

func NewRecordValueFor(
	accountID domain_model_account_id.AccountID,
	domainName domain_aggregate_domain_model_domain_name.DomainName,
) (RecordValue, error) {
	sum := sha256.Sum256([]byte(accountID.GetValue() + ":" + domainName.GetValue()))

	x, err := NewRecordValue(recordValuePrefix + hex.EncodeToString(sum[:]))
	if err != nil {
		return RecordValue{}, fmt.Errorf("new record value for account and domain: %w", err)
	}

	return x, nil
}
