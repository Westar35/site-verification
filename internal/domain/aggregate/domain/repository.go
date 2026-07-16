package domain_aggregate_domain

import (
	"context"
	"errors"

	accountID "github.com/Westar35/dns-check-service/internal/domain/model/account_id"
	normalizeddomainname "github.com/Westar35/dns-check-service/internal/domain/model/normalized_domain_name"
)

var ErrNotFound = errors.New("domain: not found")

type Repository interface {
	FindByDomain(ctx context.Context, domain normalizeddomainname.NormalizedDomainName) (Domain, error)
	Save(ctx context.Context, d Domain) error
	Delete(ctx context.Context, domain normalizeddomainname.NormalizedDomainName) error
	ListByAccount(ctx context.Context, accountID accountID.AccountID) ([]Domain, error)
}
