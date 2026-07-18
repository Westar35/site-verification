package domain_aggregate_domain_model_dns_record

import (
	"fmt"

	"github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/domain-name"
	"github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/record-name"
	"github.com/Westar35/site-verification/internal/domain/aggregate/domain/model/record-value"
	"github.com/Westar35/site-verification/internal/domain/model/account-id"
)

const recordTypeTXT = "TXT"

type DNSRecord struct {
	recordName  domain_aggregate_domain_model_record_name.RecordName
	recordType  string
	recordValue domain_aggregate_domain_model_record_value.RecordValue
}

func (x DNSRecord) GetRecordName() domain_aggregate_domain_model_record_name.RecordName {
	return x.recordName
}

func (x DNSRecord) GetRecordType() string {
	return x.recordType
}

func (x DNSRecord) GetRecordValue() domain_aggregate_domain_model_record_value.RecordValue {
	return x.recordValue
}

func NewDNSRecordFor(
	accountID domain_model_account_id.AccountID,
	domainName domain_aggregate_domain_model_domain_name.DomainName,
) (DNSRecord, error) {
	recordName, err := domain_aggregate_domain_model_record_name.NewRecordNameFor(domainName)
	if err != nil {
		return DNSRecord{}, fmt.Errorf("new dns record name: %w", err)
	}

	recordValue, err := domain_aggregate_domain_model_record_value.NewRecordValueFor(accountID, domainName)
	if err != nil {
		return DNSRecord{}, fmt.Errorf("new dns record value: %w", err)
	}

	return DNSRecord{
		recordType:  recordTypeTXT,
		recordName:  recordName,
		recordValue: recordValue,
	}, nil
}
