package domain_aggregate_domain

import (
	"errors"
	"time"

	command "github.com/Westar35/dns-check-service/internal/domain/aggregate/domain/command"
	verifieduser "github.com/Westar35/dns-check-service/internal/domain/aggregate/domain/model/verified_user"
	accountID "github.com/Westar35/dns-check-service/internal/domain/model/account_id"
	dnsrecord "github.com/Westar35/dns-check-service/internal/domain/model/dns_record"
	domainstatus "github.com/Westar35/dns-check-service/internal/domain/model/domain_status"
	domain "github.com/Westar35/dns-check-service/internal/domain/model/normalized_domain_name"
	hash "github.com/Westar35/dns-check-service/internal/domain/model/verification_hash"
)

type Domain struct {
	domainName    domain.NormalizedDomainName
	verifiedUsers []verifieduser.VerifiedUser
	createdAt     time.Time
	updatedAt     time.Time
}

func Create(cmd command.CreateDomain) (Domain, error) {
	domainName, err := domain.NewNormalizedDomainName(cmd.Domain)
	if err != nil {
		return Domain{}, err
	}
	accountID, err := accountID.New(cmd.AccountID)
	if err != nil {
		return Domain{}, err
	}
	now := time.Now().UTC()
	user, err := buildVerifiedUser(accountID, domainName, cmd.VerifiedAt, now, now)
	if err != nil {
		return Domain{}, err
	}

	d := Domain{
		domainName:    domainName,
		verifiedUsers: []verifieduser.VerifiedUser{user},
		createdAt:     now,
		updatedAt:     now,
	}
	if err := d.validate(); err != nil {
		return Domain{}, err
	}
	return d, nil
}

func buildVerifiedUser(
	accountID accountID.AccountID,
	domainName domain.NormalizedDomainName,
	verifiedAt, createdAt, updatedAt time.Time,
) (verifieduser.VerifiedUser, error) {
	hash, err := hash.GenerateVerificationHash(accountID, domainName)
	if err != nil {
		return verifieduser.VerifiedUser{}, err
	}
	record, err := dnsrecord.BuildDNSRecordFor(domainName, hash)
	if err != nil {
		return verifieduser.VerifiedUser{}, err
	}
	return verifieduser.NewVerifiedUser(accountID, hash, record, verifiedAt, createdAt, updatedAt)
}

func (d Domain) AddOrUpdateVerifiedUser(cmd command.AddOrUpdateVerifiedUser) (Domain, error) {
	accountID, err := accountID.New(cmd.AccountID)
	if err != nil {
		return d, err
	}

	now := time.Now().UTC()
	createdAt := now
	idx := -1
	for i, u := range d.verifiedUsers {
		if u.AccountID() == accountID {
			idx = i
			createdAt = u.CreatedAt()
			break
		}
	}

	user, err := buildVerifiedUser(accountID, d.domainName, cmd.VerifiedAt, createdAt, now)
	if err != nil {
		return d, err
	}

	users := append([]verifieduser.VerifiedUser{}, d.verifiedUsers...)
	if idx >= 0 {
		users[idx] = user
	} else {
		users = append(users, user)
	}

	d.verifiedUsers = users
	d.updatedAt = now
	return d, nil
}

func (d Domain) RemoveVerifiedUser(cmd command.RemoveVerifiedUser) (Domain, error) {
	accountID, err := accountID.New(cmd.AccountID)
	if err != nil {
		return d, err
	}

	users := make([]verifieduser.VerifiedUser, 0, len(d.verifiedUsers))
	for _, u := range d.verifiedUsers {
		if u.AccountID() != accountID {
			users = append(users, u)
		}
	}

	d.verifiedUsers = users
	d.updatedAt = time.Now().UTC()
	return d, nil
}

func (d Domain) VerifiedUserFor(accountID accountID.AccountID) (verifieduser.VerifiedUser, bool) {
	for _, u := range d.verifiedUsers {
		if u.AccountID() == accountID {
			return u, true
		}
	}
	return verifieduser.VerifiedUser{}, false
}

func (d Domain) StatusFor(accountID accountID.AccountID) domainstatus.DomainStatus {
	if _, ok := d.VerifiedUserFor(accountID); ok {
		return domainstatus.StatusVerified
	}
	return domainstatus.StatusNotVerified
}

func (d Domain) validate() error {
	if d.createdAt.IsZero() || d.updatedAt.IsZero() {
		return errors.New("domain: timestamps required")
	}
	return nil
}

func (d Domain) IsEmpty() bool { return len(d.verifiedUsers) == 0 }
