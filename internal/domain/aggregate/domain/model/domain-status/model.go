package domain_aggregate_domain_model_domain_status

var (
	StatusNotVerified = DomainStatus{value: "NOT_VERIFIED"}
	StatusVerified    = DomainStatus{value: "VERIFIED"}
)

type DomainStatus struct {
	value string
}

func (x DomainStatus) GetValue() string {
	return x.value
}
