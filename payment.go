package helpers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/go-zoox/fetch"
)

var (
	marchant = os.Getenv("PAYMENT_ID")
	apikey   = os.Getenv("PAYMENT_APIKEY")
	dateTime = time.Now().Format("2006-01-02 15:04:05")
)

// type Payment struct {
// 	MARCHANT string "json:merchantcode"
// 	AMOUNT   string "json:amount"
// 	DATETIME string "json:datetime"
// 	SIGNATUR string "json:signature"
// }

func GetPayment(amount string) *bytes.Buffer {
	hash := sha256.Sum256([]byte(marchant + amount + dateTime + apikey))
	res, err := fetch.Post(os.Getenv("PAYMENT_URL")+"/webapi/api/merchant/paymentmethod/getpaymentmethod",
		&fetch.Config{Body: map[string]string{
			"merchantcode": marchant,
			"amount":       amount,
			"datetime":     dateTime,
			"signature":    hex.EncodeToString(hash[:]),
		}})

	if err != nil {
		fmt.Println(err)
	}

	return bytes.NewBuffer(res.Body)
}

func TransaksiPayment(orderId string, amount string, paymentCode string, customerNumber string) *bytes.Buffer {
	hash := sha256.Sum256([]byte(marchant + orderId + apikey))
	res, err := fetch.Post(os.Getenv("PAYMENT_URL")+"/webapi/api/merchant/v2/inquiry",
		&fetch.Config{Body: map[string]string{
			"merchantcode":    marchant,
			"merchantOrderId": dateTime,
			"signature":       hex.EncodeToString(hash[:]),
		}})

	if err != nil {
		fmt.Println(err)
	}

	return bytes.NewBuffer(res.Body)
}

func TransaksiCheck(orderId string) *bytes.Buffer {
	hash := sha256.Sum256([]byte(marchant + orderId + apikey))
	res, err := fetch.Post(os.Getenv("PAYMENT_URL")+"/webapi/api/merchant/transactionStatus",
		&fetch.Config{Body: map[string]string{
			"merchantcode":    marchant,
			"merchantOrderId": dateTime,
			"signature":       hex.EncodeToString(hash[:]),
		}})

	if err != nil {
		fmt.Println(err)
	}

	return bytes.NewBuffer(res.Body)
}
