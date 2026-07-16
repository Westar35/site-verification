package domain_aggregate_domain

import (
	"time"

	verifieduser "github.com/Westar35/dns-check-service/internal/domain/aggregate/domain/model/verified_user"
	domain "github.com/Westar35/dns-check-service/internal/domain/model/normalized_domain_name"
)

func Restore(
	domainName domain.NormalizedDomainName,
	verifiedUsers []verifieduser.VerifiedUser,
	createdAt, updatedAt time.Time,
) (Domain, error) {
	d := Domain{
		domainName:    domainName,
		verifiedUsers: verifiedUsers,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}
	if err := d.validate(); err != nil {
		return Domain{}, err
	}
	return d, nil
}

func (d Domain) DomainName() domain.NormalizedDomainName    { return d.domainName }
func (d Domain) VerifiedUsers() []verifieduser.VerifiedUser { return d.verifiedUsers }
func (d Domain) CreatedAt() time.Time                       { return d.createdAt }
func (d Domain) UpdatedAt() time.Time                       { return d.updatedAt }
