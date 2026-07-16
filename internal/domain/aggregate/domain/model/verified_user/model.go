package domain_aggregate_domain_model_verified_user

import (
	"errors"
	"time"

	accountID "github.com/Westar35/dns-check-service/internal/domain/model/account_id"
	dnsrecord "github.com/Westar35/dns-check-service/internal/domain/model/dns_record"
	varificationHash "github.com/Westar35/dns-check-service/internal/domain/model/verification_hash"
)

type VerifiedUser struct {
	accountID        accountID.AccountID
	verificationHash varificationHash.VerificationHash
	record           dnsrecord.DNSRecord
	verifiedAt       time.Time
	createdAt        time.Time
	updatedAt        time.Time
}

func NewVerifiedUser(
	accountID accountID.AccountID,
	verificationHash varificationHash.VerificationHash,
	record dnsrecord.DNSRecord,
	verifiedAt, createdAt, updatedAt time.Time,
) (VerifiedUser, error) {
	u := VerifiedUser{
		accountID:        accountID,
		verificationHash: verificationHash,
		record:           record,
		verifiedAt:       verifiedAt,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
	}
	if err := u.validate(); err != nil {
		return VerifiedUser{}, err
	}
	return u, nil
}

func (u VerifiedUser) validate() error {
	if u.verifiedAt.IsZero() || u.createdAt.IsZero() || u.updatedAt.IsZero() {
		return errors.New("verified user: timestamps required")
	}
	return nil
}

func (u VerifiedUser) AccountID() accountID.AccountID                      { return u.accountID }
func (u VerifiedUser) VerificationHash() varificationHash.VerificationHash { return u.verificationHash }
func (u VerifiedUser) Record() dnsrecord.DNSRecord                         { return u.record }
func (u VerifiedUser) VerifiedAt() time.Time                               { return u.verifiedAt }
func (u VerifiedUser) CreatedAt() time.Time                                { return u.createdAt }
func (u VerifiedUser) UpdatedAt() time.Time                                { return u.updatedAt }
