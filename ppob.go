package helpers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"

	models "github.com/Payphone-Digital/Models/ppob"
	"github.com/go-zoox/fetch"
)

func CheckSaldo() *bytes.Buffer {
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + "depo"))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/cek-saldo", &fetch.Config{
		Body: models.GlobalAuth{
			CMD:      "deposit",
			USERNAME: os.Getenv("PPOB_USERNAME"),
			SIGN:     hex.EncodeToString(hash[:]),
		},
	})
	if err != nil {
		panic(err)
	}
	jsonValue, _ := json.Marshal(models.SaldoPpob{Saldo: res.Get("data").Get("deposit").Int()})
	return bytes.NewBuffer(jsonValue)
}

func DepositSaldo(at int64, bk string, or string) *bytes.Buffer {
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + "deposit"))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/deposit", &fetch.Config{
		Body: models.SaldoDepositAuth{
			USERNAME: os.Getenv("PPOB_USERNAME"),
			SIGN:     hex.EncodeToString(hash[:]),
			AMOUNT:   at,
			BANK:     bk,
			OWNER:    or,
		},
	})
	if err != nil {
		panic(err)
	}
	jsonValue, _ := json.Marshal(models.ResponseSaldoDeposit{
		RC:     res.Get("data").Get("rc").String(),
		AMOUNT: res.Get("data").Get("amount").Int(),
		NOTES:  res.Get("data").Get("notes").String(),
	})
	return bytes.NewBuffer(jsonValue)
}

func PraBayar() *bytes.Buffer {
	var hargas models.PraBayars
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + "pricelist"))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/price-list", &fetch.Config{
		Body: models.GlobalAuth{
			CMD:      "deposit",
			USERNAME: os.Getenv("PPOB_USERNAME"),
			SIGN:     hex.EncodeToString(hash[:]),
		},
	})
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(res.Get("data").String()), &hargas)
	jsonValue, _ := json.Marshal(hargas)
	return bytes.NewBuffer(jsonValue)
}

func TransaksiPraBayar(ref string, sku string, cusno string) *bytes.Buffer {
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + ref))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/transaction", &fetch.Config{
		Body: models.TransaksiPraBayar{
			USERNAME:       os.Getenv("PPOB_USERNAME"),
			SIGN:           hex.EncodeToString(hash[:]),
			BUYER_SKU_CODE: sku,
			REF_ID:         ref,
			COSTUMER_NO:    cusno,
		},
	})
	if err != nil {
		panic(err)
	}
	jsonValue, _ := json.Marshal(models.ResponsePraBayar{
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
	return bytes.NewBuffer(jsonValue)
}

func Pascabayar(ref string, sku string, cusno string) *bytes.Buffer {
	var harga models.ResponsePasca
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + ref))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/transaction", &fetch.Config{
		Body: models.CheckStatus{
			USERNAME:       os.Getenv("PPOB_USERNAME"),
			SIGN:           hex.EncodeToString(hash[:]),
			COMMANDS:       "inq-pasca",
			REF_ID:         ref,
			BUYER_SKU_CODE: sku,
			COSTUMER_NO:    cusno,
		},
	})
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(res.Get("data").String()), &harga)
	jsonValue, _ := json.Marshal(harga)
	return bytes.NewBuffer(jsonValue)
}

func TransaksiPascabayar(ref string, sku string, cusno string) *bytes.Buffer {
	var harga models.ResponsePasca
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + ref))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/transaction", &fetch.Config{
		Body: models.CheckStatus{
			USERNAME:       os.Getenv("PPOB_USERNAME"),
			SIGN:           hex.EncodeToString(hash[:]),
			COMMANDS:       "inq-pasca",
			REF_ID:         ref,
			BUYER_SKU_CODE: sku,
			COSTUMER_NO:    cusno,
		},
	})
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(res.Get("data").String()), &harga)
	jsonValue, _ := json.Marshal(harga)
	return bytes.NewBuffer(jsonValue)
}

func InquirePln(cusno string) *bytes.Buffer {
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/transaction", &fetch.Config{
		Body: models.InquirePLN{
			COMMANDS:    "pln-subscribe",
			COSTUMER_NO: cusno,
		},
	})

	if err != nil {
		panic(err)
	}
	jsonValue, _ := json.Marshal(models.ResponseInquirePLN{
		COSTUMER_NO:   res.Get("data").Get("customer_no").String(),
		METER_NO:      res.Get("data").Get("meter_no").String(),
		SUBSCRIBER_ID: res.Get("data").Get("subscriber_id").String(),
		NAME:          res.Get("data").Get("name").String(),
		SEGMENT_POWER: res.Get("data").Get("segment_power").String(),
	})
	return bytes.NewBuffer(jsonValue)
}

func CheckTagihan(ref string, sku string, cusno string) *bytes.Buffer {
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + ref))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/transaction", &fetch.Config{
		Body: models.CheckStatus{
			USERNAME:       os.Getenv("PPOB_USERNAME"),
			SIGN:           hex.EncodeToString(hash[:]),
			COMMANDS:       "status-pasca",
			REF_ID:         ref,
			BUYER_SKU_CODE: sku,
			COSTUMER_NO:    cusno,
		},
	})
	if err != nil {
		panic(err)
	}
	jsonValue, _ := json.Marshal(models.ResponseCheckStatus{
		BUYER_SKU_CODE: res.Get("data").Get("buyer_sku_code").String(),
		COSTUMER_NO:    res.Get("data").Get("customer_no").String(),
		MESSAGE:        res.Get("data").Get("message").String(),
		REF_ID:         res.Get("data").Get("ref_id").String(),
		SN:             res.Get("data").Get("sn").String(),
		STATUS:         res.Get("data").Get("status").String(),
	})
	return bytes.NewBuffer(jsonValue)
}
