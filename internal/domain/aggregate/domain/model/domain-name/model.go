package domain_aggregate_domain_model_domain_name

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	domainPattern   = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)
	ErrInvalidModel = errors.New("invalid model")
)

type DomainName struct {
	value string
}

func (x DomainName) GetValue() string {
	return x.value
}

func (x DomainName) Validate() error {
	if !domainPattern.MatchString(x.value) {
		return ErrInvalidModel
	}

	return nil
}

func NewDomainName(raw string) (DomainName, error) {
	x := DomainName{value: normalize(raw)}
	if err := x.Validate(); err != nil {
		return DomainName{}, fmt.Errorf("new domain name: %w", err)
	}

	return x, nil
}

func normalize(value string) string {
	value = stripScheme(value)
	value = stripPathQueryFragment(value)
	value = stripPort(value)
	value = strings.TrimSuffix(value, ".")
	value = strings.ToLower(strings.TrimSpace(value))

	return value
}

func stripPathQueryFragment(value string) string {
	if idx := strings.IndexAny(value, "/?#"); idx != -1 {
		return value[:idx]
	}

	return value
}

func stripPort(value string) string {
	if idx := strings.LastIndex(value, ":"); idx != -1 {
		return value[:idx]
	}

	return value
}

func stripScheme(value string) string {
	lower := strings.ToLower(value)

	switch {
	case strings.HasPrefix(lower, "https://"):
		return value[len("https://"):]
	case strings.HasPrefix(lower, "http://"):
		return value[len("http://"):]
	default:
		return value
	}
}
