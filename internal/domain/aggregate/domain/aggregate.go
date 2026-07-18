package domain_aggregate_domain

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/domain-name"
	"github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/verified-user"
)

var errInvalidAggregate = errors.New("invalid aggregate")

type Domain struct {
	domainName    domain_aggregate_domain_model_domain_name.DomainName
	verifiedUsers []domain_aggregate_domain_model_verified_user.VerifiedUser
	createdAt     time.Time
	updatedAt     time.Time
}

type Option func(x *Domain)

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
		return errInvalidAggregate
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
