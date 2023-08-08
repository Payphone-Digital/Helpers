package helpers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/go-zoox/fetch"
)

func Ppob(url, username, apikey string) *PpobNeed {
	return &PpobNeed{
		Url:      url,
		Username: username,
		Apikey:   apikey,
	}
}

func (p *PpobNeed) CheckSaldoPpob() (*bytes.Buffer, error) {
	hash := md5.Sum([]byte(p.Username + p.Apikey + "depo"))
	res, err := fetch.Post(p.Url+"/v1/cek-saldo", &fetch.Config{
		Body: GlobalAuth{
			CMD:      "deposit",
			USERNAME: p.Username,
			SIGN:     hex.EncodeToString(hash[:]),
		},
	})
	jsonValue, _ := json.Marshal(SaldoPpob{Saldo: res.Get("data").Get("deposit").Int()})
	return bytes.NewBuffer(jsonValue), err
}

func (p *PpobNeed) DepositSaldoPpob(at int64, bk, or string) (*bytes.Buffer, error) {
	hash := md5.Sum([]byte(p.Username + p.Apikey + "deposit"))
	res, err := fetch.Post(p.Url+"/v1/deposit", &fetch.Config{
		Body: SaldoDepositAuth{
			USERNAME: p.Username,
			SIGN:     hex.EncodeToString(hash[:]),
			AMOUNT:   at,
			BANK:     bk,
			OWNER:    or,
		},
	})
	jsonValue, _ := json.Marshal(ResponseSaldoDeposit{
		RC:     res.Get("data").Get("rc").String(),
		AMOUNT: res.Get("data").Get("amount").Int(),
		NOTES:  res.Get("data").Get("notes").String(),
	})
	return bytes.NewBuffer(jsonValue), err
}

func (p *PpobNeed) PraBayarPpob() (*bytes.Buffer, error) {
	var hargas PraBayars
	hash := md5.Sum([]byte(p.Username + p.Apikey + "pricelist"))
	res, err := fetch.Post(p.Url+"/v1/price-list", &fetch.Config{
		Body: GlobalAuth{
			CMD:      "deposit",
			USERNAME: p.Username,
			SIGN:     hex.EncodeToString(hash[:]),
		},
	})
	json.Unmarshal([]byte(res.Get("data").String()), &hargas)
	jsonValue, _ := json.Marshal(hargas)
	return bytes.NewBuffer(jsonValue), err
}

func (p *PpobNeed) TransaksiPraBayarPpob(ref, sku, cusno string) (*bytes.Buffer, error) {
	hash := md5.Sum([]byte(p.Username + p.Apikey + ref))
	res, err := fetch.Post(p.Url+"/v1/transaction", &fetch.Config{
		Body: TransaksiPraBayar{
			USERNAME:       p.Username,
			SIGN:           hex.EncodeToString(hash[:]),
			BUYER_SKU_CODE: sku,
			REF_ID:         ref,
			COSTUMER_NO:    cusno,
		},
	})
	jsonValue, _ := json.Marshal(ResponsePraBayar{
		BUYYER_LAST_SALDO: res.Get("data").Get("buyer_last_saldo").Int(),
		BUYER_SKU_CODE:    res.Get("data").Get("buyer_sku_code").String(),
		COSTUMER_NO:       res.Get("data").Get("customer_no").String(),
		MESSAGE:           res.Get("data").Get("message").String(),
		PRICE:             res.Get("data").Get("price").Int(),
		RC:                res.Get("data").Get("rc").String(),
		REF_ID:            res.Get("data").Get("ref_id").String(),
		SN:                res.Get("data").Get("sn").String(),
		STATUS:            res.Get("data").Get("status").String(),
		TELE:              res.Get("data").Get("tele").Int(),
		WA:                res.Get("data").Get("wa").Int(),
	})
	return bytes.NewBuffer(jsonValue), err
}

func (p *PpobNeed) PascabayarPpob(ref, sku, cusno string) (*bytes.Buffer, error) {
	var harga ResponsePasca
	hash := md5.Sum([]byte(p.Username + p.Apikey + ref))
	res, err := fetch.Post(p.Url+"/v1/transaction", &fetch.Config{
		Body: CheckStatus{
			USERNAME:       p.Username,
			SIGN:           hex.EncodeToString(hash[:]),
			COMMANDS:       "inq-pasca",
			REF_ID:         ref,
			BUYER_SKU_CODE: sku,
			COSTUMER_NO:    cusno,
		},
	})
	json.Unmarshal([]byte(res.Get("data").String()), &harga)
	jsonValue, _ := json.Marshal(harga)
	return bytes.NewBuffer(jsonValue), err
}

func (p *PpobNeed) TransaksiPascabayarPpob(ref, sku, cusno string) (*bytes.Buffer, error) {
	var harga ResponsePasca
	hash := md5.Sum([]byte(p.Username + p.Apikey + ref))
	res, err := fetch.Post(p.Url+"/v1/transaction", &fetch.Config{
		Body: CheckStatus{
			USERNAME:       p.Username,
			SIGN:           hex.EncodeToString(hash[:]),
			COMMANDS:       "inq-pasca",
			REF_ID:         ref,
			BUYER_SKU_CODE: sku,
			COSTUMER_NO:    cusno,
		},
	})
	json.Unmarshal([]byte(res.Get("data").String()), &harga)
	jsonValue, _ := json.Marshal(harga)
	return bytes.NewBuffer(jsonValue), err
}

func (p *PpobNeed) InquirePlnPpob(cusno string) (*bytes.Buffer, error) {
	res, err := fetch.Post(p.Url+"/v1/transaction", &fetch.Config{
		Body: InquirePLN{
			COMMANDS:    "pln-subscribe",
			COSTUMER_NO: cusno,
		},
	})
	jsonValue, _ := json.Marshal(ResponseInquirePLN{
		COSTUMER_NO:   res.Get("data").Get("customer_no").String(),
		METER_NO:      res.Get("data").Get("meter_no").String(),
		SUBSCRIBER_ID: res.Get("data").Get("subscriber_id").String(),
		NAME:          res.Get("data").Get("name").String(),
		SEGMENT_POWER: res.Get("data").Get("segment_power").String(),
	})
	return bytes.NewBuffer(jsonValue), err
}

func (p *PpobNeed) CheckTagihanPpob(ref, sku, cusno string) (*bytes.Buffer, error) {
	hash := md5.Sum([]byte(p.Username + p.Apikey + ref))
	res, err := fetch.Post(p.Url+"/v1/transaction", &fetch.Config{
		Body: CheckStatus{
			USERNAME:       p.Username,
			SIGN:           hex.EncodeToString(hash[:]),
			COMMANDS:       "status-pasca",
			REF_ID:         ref,
			BUYER_SKU_CODE: sku,
			COSTUMER_NO:    cusno,
		},
	})
	jsonValue, _ := json.Marshal(ResponseCheckStatus{
		BUYER_SKU_CODE: res.Get("data").Get("buyer_sku_code").String(),
		COSTUMER_NO:    res.Get("data").Get("customer_no").String(),
		MESSAGE:        res.Get("data").Get("message").String(),
		REF_ID:         res.Get("data").Get("ref_id").String(),
		SN:             res.Get("data").Get("sn").String(),
		STATUS:         res.Get("data").Get("status").String(),
	})
	return bytes.NewBuffer(jsonValue), err
}
