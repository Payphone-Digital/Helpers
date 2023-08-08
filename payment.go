package helpers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/go-zoox/fetch"
)

func Payment(url, merchant, apikey string) *PaymentNeed {
	return &PaymentNeed{
		Url:      url,
		Merchant: merchant,
		Apikey:   apikey,
		Times:    time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (p *PaymentNeed) Method(amount string) (*bytes.Buffer, error) {
	hash := sha256.Sum256([]byte(p.Merchant + amount + p.Times + p.Apikey))
	res, err := fetch.Post(p.Url+"/webapi/api/merchant/paymentmethod/getpaymentmethod",
		&fetch.Config{Body: map[string]string{
			"merchantcode": p.Merchant,
			"amount":       amount,
			"datetime":     p.Times,
			"signature":    hex.EncodeToString(hash[:]),
		}})
	return bytes.NewBuffer(res.Body), err
}

func (p *PaymentNeed) Transaksi(orderId, amount, paymentCode, customerNumber string) (*bytes.Buffer, error) {
	hash := sha256.Sum256([]byte(p.Merchant + orderId + p.Apikey))
	res, err := fetch.Post(p.Url+"/webapi/api/merchant/v2/inquiry",
		&fetch.Config{Body: map[string]string{
			"merchantcode":    p.Merchant,
			"merchantOrderId": p.Times,
			"signature":       hex.EncodeToString(hash[:]),
		}})
	return bytes.NewBuffer(res.Body), err
}

func (p *PaymentNeed) TransaksiCheck(orderId string) (*bytes.Buffer, error) {
	hash := sha256.Sum256([]byte(p.Merchant + orderId + p.Apikey))
	res, err := fetch.Post(p.Url+"/webapi/api/merchant/transactionStatus",
		&fetch.Config{Body: map[string]string{
			"merchantcode":    p.Merchant,
			"merchantOrderId": p.Times,
			"signature":       hex.EncodeToString(hash[:]),
		}})
	return bytes.NewBuffer(res.Body), err
}
