package domain_model_normalized_domain_name

import (
	"errors"
	"regexp"
	"strings"

	domain_name "github.com/Westar35/dns-check-service/internal/domain/model/domain_name"
)

var domainPattern = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

type NormalizedDomainName struct {
	value string
}

func Normalize(raw domain_name.DomainName) (NormalizedDomainName, error) {
	v := raw.String()
	v = stripScheme(v)
	v = stripPathQueryFragment(v)
	v = stripPort(v)
	v = strings.TrimSuffix(v, ".")
	v = strings.ToLower(strings.TrimSpace(v))

	return NewNormalizedDomainName(v)
}

func NewNormalizedDomainName(raw string) (NormalizedDomainName, error) {
	n := NormalizedDomainName{value: raw}
	if err := n.validate(); err != nil {
		return NormalizedDomainName{}, err
	}
	return n, nil
}

func (m NormalizedDomainName) validate() error {
	if !domainPattern.MatchString(m.value) {
		return errors.New("normalized domain name: invalid format")
	}
	return nil
}

func (m NormalizedDomainName) String() string {
	return m.value
}

func stripScheme(v string) string {
	lower := strings.ToLower(v)
	switch {
	case strings.HasPrefix(lower, "https://"):
		return v[len("https://"):]
	case strings.HasPrefix(lower, "http://"):
		return v[len("http://"):]
	default:
		return v
	}
}

func stripPathQueryFragment(v string) string {
	if idx := strings.IndexAny(v, "/?#"); idx != -1 {
		return v[:idx]
	}
	return v
}

func stripPort(v string) string {
	if idx := strings.LastIndex(v, ":"); idx != -1 {
		return v[:idx]
	}
	return v
}
