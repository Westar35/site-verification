package domain_model_verification_hash

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	accountID "github.com/Westar35/dns-check-service/internal/domain/model/account_id"
	domain "github.com/Westar35/dns-check-service/internal/domain/model/normalized_domain_name"
)

type VerificationHash struct {
	value string
}

func (m VerificationHash) String() string {
	return m.value
}

func GenerateVerificationHash(accountID accountID.AccountID, domain domain.NormalizedDomainName) (VerificationHash, error) {
	sum := sha256.Sum256([]byte(accountID.String() + ":" + domain.String()))
	return NewVerificationHash(hex.EncodeToString(sum[:]))
}

func NewVerificationHash(raw string) (VerificationHash, error) {
	h := VerificationHash{value: raw}
	if err := h.validate(); err != nil {
		return VerificationHash{}, err
	}
	return h, nil
}

func (h VerificationHash) validate() error {
	if len(h.value) != 64 {
		return errors.New("VerificationHash: invalid length")
	}
	return nil
}
