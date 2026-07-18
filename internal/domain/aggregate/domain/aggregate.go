package domain_aggregate_domain

import (
	"errors"
	"fmt"
	"slices"
	"time"

	domain_aggregate_domain_model_domain_name "github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/domain-name"
	domain_aggregate_domain_model_verified_user "github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/verified-user"
)

type Domain struct {
	createdAt     time.Time
	domainName    domain_aggregate_domain_model_domain_name.DomainName
	updatedAt     time.Time
	verifiedUsers []domain_aggregate_domain_model_verified_user.VerifiedUser
}

func (x Domain) GetCreatedAt() time.Time {
	return x.createdAt
}

func (x Domain) GetDomainName() domain_aggregate_domain_model_domain_name.DomainName {
	return x.domainName
}

func (x Domain) GetUpdatedAt() time.Time {
	return x.updatedAt
}

func (x Domain) GetVerifiedUsers() []domain_aggregate_domain_model_verified_user.VerifiedUser {
	return slices.Clone(x.verifiedUsers)
}

func (x Domain) Validate() error {
	if x.domainName.GetValue() == "" {
		return errors.New("invalid aggregate")
	}

	return nil
}

func NewDomain(
	domainName domain_aggregate_domain_model_domain_name.DomainName,
	opts ...Option,
) (*Domain, error) {
	x := &Domain{domainName: domainName}

	for i := range opts {
		opts[i](x)
	}

	if err := x.Validate(); err != nil {
		return nil, fmt.Errorf("new domain: %w", err)
	}

	return x, nil
}

type Option func(x *Domain)

func WithCreatedAt(createdAt time.Time) Option {
	return func(x *Domain) {
		x.createdAt = createdAt
	}
}

func WithUpdatedAt(updatedAt time.Time) Option {
	return func(x *Domain) {
		x.updatedAt = updatedAt
	}
}

func WithVerifiedUsers(verifiedUsers []domain_aggregate_domain_model_verified_user.VerifiedUser) Option {
	return func(x *Domain) {
		x.verifiedUsers = slices.Clone(verifiedUsers)
	}
}
