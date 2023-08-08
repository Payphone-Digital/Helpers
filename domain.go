package helpers

import (
	"bytes"

	"github.com/go-zoox/fetch"
)

func (d *DomainNeed) Domain(url string, header map[string]string) *DomainNeed {
	return &DomainNeed{
		Url:    url,
		Header: header,
	}
}

//Start Account Domain

func (d *DomainNeed) BalanceAccount() (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/account/balance",
		&fetch.Config{Headers: d.Header})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PriceAccount() (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/account/prices",
		&fetch.Config{Headers: d.Header})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) TransactionAccount(data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/account/transactions",
		&fetch.Config{Headers: d.Header, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Account Domain

// Start Costumer
func (d *DomainNeed) GetCostumer(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/customers",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostCostumer(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/customers",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Costumer

//Start Domain Forwarding

func (d *DomainNeed) GetForwardingDomain(id string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/domain_forwarding",
		&fetch.Config{Headers: d.Header})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) UpdateForwardingDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/domain_forwarding",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Domain Forwarding

//Start Privacy Protection

func (d *DomainNeed) GetPrivacyProtection(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/privacy_protection",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostPrivacyProtection(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/privacy_protection/buy",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PutPrivacyProtection(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/privacy_protection",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DeletePrivacyProtection(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(d.Url+"/v1/domains/"+id+"/privacy_protection",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Privacy Protection

// Start Domain
// Start Search Domain
func (d *DomainNeed) GetSearchBuyDomain(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/suggestion",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) GetAllDomain(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) GetByNameDomain(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/details-by-name",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End Search Domain
// Start Register Domain
func (d *DomainNeed) RegisterDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Register Domain

// Start Transfer Domain

func (d *DomainNeed) TransferDomain(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/transfer",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) ValidityTransferDomain(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/transfer/validity",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) CencelTransferDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/transfer/cancel",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) ResendTransferDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/transfer/resend_approval_email",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Transfer Domain

//Start Code Epp

func (d *DomainNeed) GetEppDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/auth_code",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) UpdateEppDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/auth_code",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Code Epp

//Start Theft Protect

func (d *DomainNeed) GetTheftProtectDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/theft_protection",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) EnableTheftProtectDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/theft_protection",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DisableTheftProtectDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(d.Url+"/v1/domains/"+id+"/theft_protection",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Theft Protect

//Start Suspend

func (d *DomainNeed) GetSuspendDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/suspended",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) EnableSuspendDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/suspended",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DisableSuspendDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(d.Url+"/v1/domains/"+id+"/suspended",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Suspend

//Start Locked

func (d *DomainNeed) GetLockedDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/locked",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) EnableLockedDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/locked",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DisableLockedDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(d.Url+"/v1/domains/"+id+"/locked",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Locked

// Start Ns
func (d *DomainNeed) GetNsDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/ns",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) UpdateNsDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/ns",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Ns

// Start Raa Verification
func (d *DomainNeed) GetRaaVerifyDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/raa_verification",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostRaaVerifyDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/raa_verification/resend",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Raa Verification

// Start Renew Domain
func (d *DomainNeed) RenewDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/renew",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End Renew Domain

// Start Restore Domain
func (d *DomainNeed) RestoreDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/restore",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End Restore Domain

// Start Move Domain
func (d *DomainNeed) MoveDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/move",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End Move Domain

// Start Irtp Verification Domain
func (d *DomainNeed) IrtpVerifyDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/irtp_verification/resend",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End Irtp Verification Domain

// Start DNSSEC
func (d *DomainNeed) GetDnssecDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/dnssec",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostDnssecDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/dnssec",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DeleteDnssecDomain(id, keytag, algorithm, digesttype, digest string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/dnssec/"+keytag+"/"+algorithm+"/"+digesttype+"/"+digest,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End DNSSEC

// Start Child DNS
func (d *DomainNeed) GetChilddnsDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/childns",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostChilddnsDomain(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/childns",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) UpdateChilddnsDomain(id, old_hosname, old_ip_address string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/childns/"+old_hosname+"/"+old_ip_address,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DeleteChilddnsDomain(id string, hosname string, ip_address string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/childns/"+hosname+"/"+ip_address,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

//End Child DNS
//End Domain

// Start Domain DNS
// Start CNAME
func (d *DomainNeed) GetDnsCname(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/dns/cname",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostDnsCname(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/dns/cname",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PutDnsCname(id, old_hostname, old_value string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/dns/cname/"+old_hostname+"/"+old_value,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DeleteDnsCname(id, hostname, value string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(d.Url+"/v1/domains/"+id+"/dns/cname/"+hostname+"/"+value,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End CNAME
// Start IPv4
func (d *DomainNeed) GetDnsIpv4(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/dns/ip",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostDnsIpv4(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/dns/ip",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PutDnsIpv4(id, old_hostname, old_value string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/dns/ip/"+old_hostname+"/"+old_value,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DeleteDnsIpv4(id, hostname, value string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(d.Url+"/v1/domains/"+id+"/dns/ip/"+hostname+"/"+value,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End IPv4
// Start IPv6
func (d *DomainNeed) GetDnsIpv6(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/dns/ipv6",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostDnsIpv6(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/dns/ipv6",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PutDnsIpv6(id, old_hostname, old_value string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/dns/ipv6/"+old_hostname+"/"+old_value,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DeleteDnsIpv6(id, hostname, value string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(d.Url+"/v1/domains/"+id+"/dns/ipv6/"+hostname+"/"+value,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End IPv6
// Start MX
func (d *DomainNeed) GetDnsMx(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/dns/mx",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostDnsMx(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/dns/mx",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PutDnsMx(id, old_hostname, old_value string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/dns/mx/"+old_hostname+"/"+old_value,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DeleteDnsMx(id, hostname, value string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(d.Url+"/v1/domains/"+id+"/dns/mx/"+hostname+"/"+value,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End MX
// Start SRV
func (d *DomainNeed) GetDnsSrv(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/dns/srv",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostDnsSrv(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/dns/srv",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PutDnsSrv(id, old_hostname, old_value, old_port, old_weight, old_priority string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/dns/srv/"+old_hostname+"/"+old_value+"/"+old_port+"/"+old_weight+"/"+old_priority,
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DeleteDnsSrv(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(d.Url+"/v1/domains/"+id+"/dns/srv",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End SRV
// Start TXT
func (d *DomainNeed) GetDnsTxt(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(d.Url+"/v1/domains/"+id+"/dns/txt",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PostDnsTxt(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(d.Url+"/v1/domains/"+id+"/dns/txt",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) PutDnsTxt(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(d.Url+"/v1/domains/"+id+"/dns/txt",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (d *DomainNeed) DeleteDnsTxt(id string, param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(d.Url+"/v1/domains/"+id+"/dns/txt",
		&fetch.Config{Headers: d.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

// End TXT
//End Domain DNS
