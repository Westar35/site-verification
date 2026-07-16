package domain_aggregate_domain_command

import (
	"time"
)

type AddOrUpdateVerifiedUser struct {
	AccountID  string
	VerifiedAt time.Time
}
