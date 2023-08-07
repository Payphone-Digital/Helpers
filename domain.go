package helpers

import (
	"bytes"
	"encoding/base64"
	"os"

	"github.com/go-zoox/fetch"
)

var basic = map[string]string{"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(os.Getenv("DOMAIN_RESID")+":"+os.Getenv("DOMAIN_APIKEY")))}

//Start Account Domain

func BalanceAccount() (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/account/balance",
		&fetch.Config{Headers: basic})
	return bytes.NewBuffer(res.Body), err
}

func PriceAccount() (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/account/prices",
		&fetch.Config{Headers: basic})
	return bytes.NewBuffer(res.Body), err
}

func TransactionAccount(data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/account/transactions",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Account Domain

// Start Costumer
func GetCostumer(data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/customers",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostCostumer(data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/customers",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Costumer

//Start Domain Forwarding

func GetForwardingDomain(id string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/domain_forwarding",
		&fetch.Config{Headers: basic})
	return bytes.NewBuffer(res.Body), err
}

func UpdateForwardingDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/domain_forwarding",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Domain Forwarding

//Start Privacy Protection

func GetPrivacyProtection(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/privacy_protection",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostPrivacyProtection(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/privacy_protection/buy",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PutPrivacyProtection(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/privacy_protection",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DeletePrivacyProtection(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/privacy_protection",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Privacy Protection

// Start Domain
// Start Search Domain
func GetSearchBuyDomain(data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/suggestion",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func GetAllDomain(data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func GetByNameDomain(data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/details-by-name",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End Search Domain
// Start Register Domain
func RegisterDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Register Domain

// Start Transfer Domain

func TransferDomain(data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/transfer",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func ValidityTransferDomain(data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/transfer/validity",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func CencelTransferDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/transfer/cancel",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func ResendTransferDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/transfer/resend_approval_email",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Transfer Domain

//Start Code Epp

func GetEppDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/auth_code",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func UpdateEppDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/auth_code",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Code Epp

//Start Theft Protect

func GetTheftProtectDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/theft_protection",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func EnableTheftProtectDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/theft_protection",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DisableTheftProtectDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/theft_protection",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Theft Protect

//Start Suspend

func GetSuspendDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/suspended",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func EnableSuspendDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/suspended",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DisableSuspendDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/suspended",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Suspend

//Start Locked

func GetLockedDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/locked",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func EnableLockedDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/locked",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DisableLockedDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/locked",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Locked

// Start Ns
func GetNsDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/ns",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func UpdateNsDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/ns",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Ns

// Start Raa Verification
func GetRaaVerifyDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/raa_verification",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostRaaVerifyDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/raa_verification/resend",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Raa Verification

// Start Renew Domain
func RenewDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/renew",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End Renew Domain

// Start Restore Domain
func RestoreDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/restore",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End Restore Domain

// Start Move Domain
func MoveDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/move",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End Move Domain

// Start Irtp Verification Domain
func IrtpVerifyDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/irtp_verification/resend",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End Irtp Verification Domain

// Start DNSSEC
func GetDnssecDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dnssec",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostDnssecDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dnssec",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DeleteDnssecDomain(id, keytag, algorithm, digesttype, digest string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dnssec/"+keytag+"/"+algorithm+"/"+digesttype+"/"+digest,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End DNSSEC

// Start Child DNS
func GetChilddnsDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/childns",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostChilddnsDomain(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/childns",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func UpdateChilddnsDomain(id, old_hosname, old_ip_address string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/childns/"+old_hosname+"/"+old_ip_address,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DeleteChilddnsDomain(id string, hosname string, ip_address string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/childns/"+hosname+"/"+ip_address,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

//End Child DNS
//End Domain

// Start Domain DNS
// Start CNAME
func GetDnsCname(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/cname",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostDnsCname(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/cname",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PutDnsCname(id, old_hostname, old_value string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/cname/"+old_hostname+"/"+old_value,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DeleteDnsCname(id, hostname, value string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/cname/"+hostname+"/"+value,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End CNAME
// Start IPv4
func GetDnsIpv4(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/ip",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostDnsIpv4(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/ip",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PutDnsIpv4(id, old_hostname, old_value string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/ip/"+old_hostname+"/"+old_value,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DeleteDnsIpv4(id, hostname, value string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/ip/"+hostname+"/"+value,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End IPv4
// Start IPv6
func GetDnsIpv6(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/ipv6",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostDnsIpv6(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/ipv6",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PutDnsIpv6(id, old_hostname, old_value string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/ipv6/"+old_hostname+"/"+old_value,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DeleteDnsIpv6(id, hostname, value string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/ipv6/"+hostname+"/"+value,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End IPv6
// Start MX
func GetDnsMx(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/mx",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostDnsMx(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/mx",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PutDnsMx(id, old_hostname, old_value string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/mx/"+old_hostname+"/"+old_value,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DeleteDnsMx(id, hostname, value string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/mx/"+hostname+"/"+value,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End MX
// Start SRV
func GetDnsSrv(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/srv",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostDnsSrv(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/srv",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PutDnsSrv(id, old_hostname, old_value, old_port, old_weight, old_priority string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/srv/"+old_hostname+"/"+old_value+"/"+old_port+"/"+old_weight+"/"+old_priority,
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DeleteDnsSrv(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/srv",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End SRV
// Start TXT
func GetDnsTxt(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/txt",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PostDnsTxt(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Post(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/txt",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func PutDnsTxt(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Put(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/txt",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

func DeleteDnsTxt(id string, data map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Delete(os.Getenv("DOMAIN_URL")+"/v1/domains/"+id+"/dns/txt",
		&fetch.Config{Headers: basic, Params: data})
	return bytes.NewBuffer(res.Body), err
}

// End TXT
//End Domain DNS
