package domain_aggregate_domain_command

import (
	"time"
)

type CreateDomain struct {
	Domain     string
	AccountID  string
	VerifiedAt time.Time
}
