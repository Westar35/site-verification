package domain_model_dns_record

import (
	domain "github.com/Westar35/dns-check-service/internal/domain/model/normalized_domain_name"
	recordname "github.com/Westar35/dns-check-service/internal/domain/model/record_name"
	recordvalue "github.com/Westar35/dns-check-service/internal/domain/model/record_value"
	verificationHash "github.com/Westar35/dns-check-service/internal/domain/model/verification_hash"
)

const (
	recordTypeTXT = "TXT"
)

type DNSRecord struct {
	recordType  string
	recordName  recordname.RecordName
	recordValue recordvalue.RecordValue
}

func BuildDNSRecordFor(
	domain domain.NormalizedDomainName,
	verificationhash verificationHash.VerificationHash,
) (DNSRecord, error) {
	recordName, err := recordname.NewRecordNameFor(domain)
	if err != nil {
		return DNSRecord{}, err
	}
	recordValue, err := recordvalue.NewRecordValueFor(verificationhash)
	if err != nil {
		return DNSRecord{}, err
	}

	return newDNSRecord(recordName, recordValue), nil
}

func newDNSRecord(name recordname.RecordName, value recordvalue.RecordValue) DNSRecord {
	r := DNSRecord{
		recordType:  recordTypeTXT,
		recordName:  name,
		recordValue: value,
	}
	return r
}

func (r DNSRecord) RecordType() string                   { return r.recordType }
func (r DNSRecord) RecordName() recordname.RecordName    { return r.recordName }
func (r DNSRecord) RecordValue() recordvalue.RecordValue { return r.recordValue }
