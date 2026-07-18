package domain_aggregate_domain_model_verified_user

import (
	"errors"
	"fmt"
	"time"

	"github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/dns-record"
	"github.com/Westar35/site-verification/internal/domain/model/account-id"
)

var ErrInvalidModel = errors.New("invalid model")

type VerifiedUser struct {
	accountID  domain_model_account_id.AccountID
	createdAt  time.Time
	record     domain_aggregate_domain_model_dns_record.DNSRecord
	updatedAt  time.Time
	verifiedAt time.Time
}

func (x VerifiedUser) GetAccountID() domain_model_account_id.AccountID {
	return x.accountID
}

func (x VerifiedUser) GetCreatedAt() time.Time {
	return x.createdAt
}

func (x VerifiedUser) GetRecord() domain_aggregate_domain_model_dns_record.DNSRecord {
	return x.record
}

func (x VerifiedUser) GetUpdatedAt() time.Time {
	return x.updatedAt
}

func (x VerifiedUser) GetVerifiedAt() time.Time {
	return x.verifiedAt
}

func (x VerifiedUser) Validate() error {
	if x.verifiedAt.IsZero() || x.createdAt.IsZero() || x.updatedAt.IsZero() {
		return ErrInvalidModel
	}

	return nil
}

func NewVerifiedUser(
	accountID domain_model_account_id.AccountID,
	record domain_aggregate_domain_model_dns_record.DNSRecord,
	verifiedAt, createdAt, updatedAt time.Time,
) (VerifiedUser, error) {
	x := VerifiedUser{
		accountID:  accountID,
		record:     record,
		verifiedAt: verifiedAt,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
	}

	if err := x.Validate(); err != nil {
		return VerifiedUser{}, fmt.Errorf("new verified user: %w", err)
	}

	return x, nil
}
