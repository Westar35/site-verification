package domain_model_domain_status

type DomainStatus struct {
	value string
}

var (
	StatusNotVerified = DomainStatus{value: "NOT_VERIFIED"}
	StatusVerified    = DomainStatus{value: "VERIFIED"}
)

func (m DomainStatus) String() string {
	return m.value
}
