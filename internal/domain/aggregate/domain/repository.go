package domain_aggregate_domain

import (
	"context"

	"github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/domain-name"
	"github.com/Westar35/site-verification/internal/domain/model/account-id"
)

type Repository interface {
	Delete(ctx context.Context, domainName domain_aggregate_domain_model_domain_name.DomainName) error
	Get(ctx context.Context, domainName domain_aggregate_domain_model_domain_name.DomainName) (*Domain, error)
	ListByAccount(ctx context.Context, accountID domain_model_account_id.AccountID) ([]*Domain, error)
	Save(ctx context.Context, domain *Domain) error
}
