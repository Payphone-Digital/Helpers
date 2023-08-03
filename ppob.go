package helpers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"

	"github.com/go-zoox/fetch"
)

func CheckSaldoPpob() *bytes.Buffer {
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + "depo"))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/cek-saldo", &fetch.Config{
		Body: GlobalAuth{
			CMD:      "deposit",
			USERNAME: os.Getenv("PPOB_USERNAME"),
			SIGN:     hex.EncodeToString(hash[:]),
		},
	})
	if err != nil {
		panic(err)
	}
	jsonValue, _ := json.Marshal(SaldoPpob{Saldo: res.Get("data").Get("deposit").Int()})
	return bytes.NewBuffer(jsonValue)
}

func DepositSaldoPpob(at int64, bk string, or string) *bytes.Buffer {
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + "deposit"))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/deposit", &fetch.Config{
		Body: SaldoDepositAuth{
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
	jsonValue, _ := json.Marshal(ResponseSaldoDeposit{
		RC:     res.Get("data").Get("rc").String(),
		AMOUNT: res.Get("data").Get("amount").Int(),
		NOTES:  res.Get("data").Get("notes").String(),
	})
	return bytes.NewBuffer(jsonValue)
}

func PraBayarPpob() *bytes.Buffer {
	var hargas PraBayars
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + "pricelist"))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/price-list", &fetch.Config{
		Body: GlobalAuth{
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

func TransaksiPraBayarPpob(ref string, sku string, cusno string) *bytes.Buffer {
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + ref))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/transaction", &fetch.Config{
		Body: TransaksiPraBayar{
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
	return bytes.NewBuffer(jsonValue)
}

func PascabayarPpob(ref string, sku string, cusno string) *bytes.Buffer {
	var harga ResponsePasca
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + ref))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/transaction", &fetch.Config{
		Body: CheckStatus{
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

func TransaksiPascabayarPpob(ref string, sku string, cusno string) *bytes.Buffer {
	var harga ResponsePasca
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + ref))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/transaction", &fetch.Config{
		Body: CheckStatus{
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

func InquirePlnPpob(cusno string) *bytes.Buffer {
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/transaction", &fetch.Config{
		Body: InquirePLN{
			COMMANDS:    "pln-subscribe",
			COSTUMER_NO: cusno,
		},
	})

	if err != nil {
		panic(err)
	}
	jsonValue, _ := json.Marshal(ResponseInquirePLN{
		COSTUMER_NO:   res.Get("data").Get("customer_no").String(),
		METER_NO:      res.Get("data").Get("meter_no").String(),
		SUBSCRIBER_ID: res.Get("data").Get("subscriber_id").String(),
		NAME:          res.Get("data").Get("name").String(),
		SEGMENT_POWER: res.Get("data").Get("segment_power").String(),
	})
	return bytes.NewBuffer(jsonValue)
}

func CheckTagihanPpob(ref string, sku string, cusno string) *bytes.Buffer {
	hash := md5.Sum([]byte(os.Getenv("PPOB_USERNAME") + os.Getenv("PPOB_APIKEY") + ref))
	res, err := fetch.Post(os.Getenv("PPOB_URL")+"/v1/transaction", &fetch.Config{
		Body: CheckStatus{
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
	jsonValue, _ := json.Marshal(ResponseCheckStatus{
		BUYER_SKU_CODE: res.Get("data").Get("buyer_sku_code").String(),
		COSTUMER_NO:    res.Get("data").Get("customer_no").String(),
		MESSAGE:        res.Get("data").Get("message").String(),
		REF_ID:         res.Get("data").Get("ref_id").String(),
		SN:             res.Get("data").Get("sn").String(),
		STATUS:         res.Get("data").Get("status").String(),
	})
	return bytes.NewBuffer(jsonValue)
}

// Global Auth
type GlobalAuth struct {
	CMD      string `bson:"cmd" json:"cmd"`
	USERNAME string `bson:"username" json:"username"`
	SIGN     string `bson:"sign" json:"sign"`
}

// Start Saldo

type SaldoPpob struct {
	Saldo int64 `bson:"saldo" json:"saldo"`
}

type SaldoDepositAuth struct {
	USERNAME string `bson:"username" json:"username"`
	SIGN     string `bson:"sign" json:"sign"`
	AMOUNT   int64  `bson:"amount" json:"amount"`
	BANK     string `bson:"Bank" json:"Bank"`
	OWNER    string `bson:"owner_name" json:"owner_name"`
}

type ResponseSaldoDeposit struct {
	RC     string `bson:"rc" json:"rc"`
	AMOUNT int64  `bson:"amount" json:"amount"`
	NOTES  string `bson:"notes" json:"notes"`
}

//End Saldo

// Start Prabayar
type PraBayar struct {
	PRODUCT_NAME          string `bson:"product_name" json:"product_name"`
	CATEGORY              string `bson:"category" json:"category"`
	BRAND                 string `bson:"brand" json:"brand"`
	TYPE                  string `bson:"type" json:"type"`
	SELLER_NAME           string `bson:"seller_name" json:"seller_name"`
	PRICE                 int64  `bson:"price" json:"price"`
	BUYER_SKU_CODE        string `bson:"buyer_sku_code" json:"buyer_sku_code"`
	BUYER_PRODUCT_STATUS  bool   `bson:"buyer_product_status" json:"buyer_product_status"`
	SELLER_PRODUCT_STATUS bool   `bson:"seller_product_status" json:"seller_product_status"`
	UNLIMITED_STOCK       bool   `bson:"unlimited_stock" json:"unlimited_stock"`
	STOCK                 int64  `bson:"stock" json:"stock"`
	MULTI                 bool   `bson:"multi" json:"multi"`
	START_CUT_OFF         string `bson:"start_cut_off" json:"start_cut_off"`
	END_CUT_OFF           string `bson:"end_cut_off" json:"end_cut_off"`
	DESC                  string `bson:"desc" json:"desc"`
}

type PraBayars []PraBayar

type TransaksiPraBayar struct {
	USERNAME       string `bson:"username" json:"username"`
	SIGN           string `bson:"sign" json:"sign"`
	BUYER_SKU_CODE string `bson:"buyer_sku_code" json:"buyer_sku_code"`
	REF_ID         string `bson:"ref_id" json:"ref_id"`
	COSTUMER_NO    string `bson:"customer_no" json:"customer_no"`
}

type ResponsePraBayar struct {
	REF_ID            string `bson:"ref_id" json:"ref_id"`
	COSTUMER_NO       string `bson:"customer_no" json:"customer_no"`
	BUYER_SKU_CODE    string `bson:"buyer_sku_code" json:"buyer_sku_code"`
	MESSAGE           string `bson:"message" json:"message"`
	STATUS            string `bson:"status" json:"status"`
	RC                string `bson:"rc" json:"rc"`
	SN                string `bson:"sn" json:"sn"`
	BUYYER_LAST_SALDO int64  `bson:"buyer_last_saldo" json:"buyer_last_saldo"`
	PRICE             int64  `bson:"price" json:"price"`
	TELE              int64  `bson:"tele" json:"tele"`
	WA                int64  `bson:"wa" json:"wa"`
}

// End Prabayar

// Start Pascabayar
type PascaBayar struct {
	PRODUCT_NAME          string   `bson:"product_name" json:"product_name"`
	CATEGORY              string   `bson:"category" json:"category"`
	BRAND                 string   `bson:"brand" json:"brand"`
	SELLER_NAME           string   `bson:"seller_name" json:"seller_name"`
	ADMIN                 int64    `bson:"admin" json:"admin"`
	COMMISSION            int64    `bson:"commission" json:"commission"`
	BUYER_SKU_CODE        string   `bson:"buyer_sku_code" json:"buyer_sku_code"`
	BUYER_PRODUCT_STATUS  bool     `bson:"buyer_product_status" json:"buyer_product_status"`
	SELLER_PRODUCT_STATUS bool     `bson:"seller_product_status" json:"seller_product_status"`
	DESC                  []string `bson:"desc" json:"desc"`
}

type PascaBayars []PascaBayar

type ResponsePasca struct {
	REF_ID           string    `bson:"ref_id" json:"ref_id"`
	COSTUMER_NO      string    `bson:"customer_no" json:"customer_no"`
	CUSTOMER_NAME    string    `bson:"customer_name" json:"customer_name"`
	BUYER_SKU_CODE   string    `bson:"buyer_sku_code" json:"buyer_sku_code"`
	ADMIN            int64     `bson:"admin" json:"admin"`
	MESSAGE          string    `bson:"message" json:"message"`
	STATUS           string    `bson:"status" json:"status"`
	RC               string    `bson:"rc" json:"rc"`
	BUYER_LAST_SALDO int64     `bson:"buyer_last_saldo" json:"buyer_last_saldo"`
	PRICE            int64     `bson:"price" json:"price"`
	SELLING_PRICE    int64     `bson:"selling_price" json:"selling_price"`
	SN               string    `bson:"sn" json:"sn"`
	DESC             PascaDesc `bson:"desc" json:"desc"`
}

type PascaDesc struct {
	TARIF                 string        `bson:"tarif" json:"tarif"`
	DAYA                  int64         `bson:"daya" json:"daya"`
	LEMBAR_TAGIHAN        int64         `bson:"lembar_tagihan" json:"lembar_tagihan"`
	ALAMAT                string        `bson:"alamat" json:"alamat"`
	JATUH_TEMPO           string        `bson:"jatuh_tempo" json:"jatuh_tempo"`
	JUMLAH_PESERTA        string        `bson:"jumlah_peserta" json:"jumlah_peserta"`
	ITEM_NAME             string        `bson:"item_name" json:"item_name"`
	NO_RANGKA             string        `bson:"no_rangka" json:"no_rangka"`
	NO_POL                string        `bson:"no_pol" json:"no_pol"`
	TENOR                 string        `bson:"tenor" json:"tenor"`
	TAHUN_PAJAK           string        `bson:"tahun_pajak" json:"tahun_pajak"`
	KELURAHAN             string        `bson:"kelurahan" json:"kelurahan"`
	KECAMATAN             string        `bson:"kecamatan" json:"kecamatan"`
	KODE_KAB_KOTA         string        `bson:"kode_kab_kota" json:"kode_kab_kota"`
	KAB_KOTA              string        `bson:"kab_kota" json:"kab_kota"`
	LUAS_TANAH            string        `bson:"luas_tanah" json:"luas_tanah"`
	LUAS_GEDUNG           string        `bson:"luas_gedung" json:"luas_gedung"`
	NOMOR_IDENTITAS       string        `bson:"nomor_identitas" json:"nomor_identitas"`
	NOMOR_RANGKA          string        `bson:"nomor_rangka" json:"nomor_rangka"`
	NOMOR_MESIN           string        `bson:"nomor_mesin" json:"nomor_mesin"`
	NOMOR_POLISI          string        `bson:"nomor_polisi" json:"nomor_polisi"`
	MILIK_KENAMA          string        `bson:"milik_kenama" json:"milik_kenama"`
	MEREK_KB              string        `bson:"merek_kb" json:"merek_kb"`
	MODEL_KB              string        `bson:"model_kb" json:"model_kb"`
	TAHUN_BUATAN          string        `bson:"tahun_buatan" json:"tahun_buatan"`
	TGL_AKHIR_PAJAK_BARU  string        `bson:"tgl_akhir_pajak_baru" json:"tgl_akhir_pajak_baru"`
	BIAYA_POKOK_BBN       string        `bson:"biaya_pokok_bbn" json:"biaya_pokok_bbn"`
	BIAYA_POKOK_SWD       string        `bson:"biaya_pokok_swd" json:"biaya_pokok_swd"`
	BIAYA_POKOK_PKB       string        `bson:"biaya_pokok_pkb" json:"biaya_pokok_pkb"`
	BIAYA_DENDA_SWD       string        `bson:"biaya_denda_swd" json:"biaya_denda_swd"`
	BIAYA_DENDA_BBN       string        `bson:"biaya_denda_bbn" json:"biaya_denda_bbn"`
	BIAYA_DENDA_PKB       string        `bson:"biaya_denda_pkb" json:"biaya_denda_pkb"`
	BIAYA_ADMIN_STNK      string        `bson:"biaya_admin_stnk" json:"biaya_admin_stnk"`
	BIAYA_ADMIN_TNKB      string        `bson:"biaya_admin_tnkb" json:"biaya_admin_tnkb"`
	BIAYA_PARKIR_POKOK    string        `bson:"biaya_parkir_pokok" json:"biaya_parkir_pokok"`
	BIAYA_PAJAK_PROGRESIF string        `bson:"biaya_pajak_progresif" json:"biaya_pajak_progresif"`
	KODE_IURAN            string        `bson:"kode_iuran" json:"kode_iuran"`
	KODE_PROGRAM          string        `bson:"kode_program" json:"kode_program"`
	JKK                   int64         `bson:"jkk" json:"jkk"`
	JKM                   int64         `bson:"jkm" json:"jkm"`
	JHT                   int64         `bson:"jht" json:"jht"`
	KANTOR_CABANG         string        `bson:"kantor_cabang" json:"kantor_cabang"`
	TGL_EFEKTIF           string        `bson:"tgl_efektif" json:"tgl_efektif"`
	TGL_EXPIRED           string        `bson:"tgl_expired" json:"tgl_expired"`
	JPK                   int64         `bson:"jpk" json:"jpk"`
	JPN                   int64         `bson:"jpn" json:"jpn"`
	NPP                   string        `bson:"npp" json:"npp"`
	KODE_DIVISI           string        `bson:"kode_divisi" json:"kode_divisi"`
	DETAIL                []PascaDetail `bson:"detail" json:"detail"`
}

type PascaDetail struct {
	PERIODE       string `bson:"periode" json:"periode"`
	NILAI_TAGIHAN string `bson:"nilai_tagihan" json:"nilai_tagihan"`
	DENDA         string `bson:"denda" json:"denda"`
	ADMIN         string `bson:"admin" json:"admin"`
	METER_AWAL    string `bson:"meter_awal" json:"meter_awal"`
	METER_AKHIR   string `bson:"meter_akhir" json:"meter_akhir"`
	BIAYA_LAIN    string `bson:"biaya_lain" json:"biaya_lain"`
	NO_REF        string `bson:"no_ref" json:"no_ref"`
}

//End Pascabayar

// Start Check Status

type CheckStatus struct {
	COMMANDS       string `bson:"commands" json:"commands"`
	USERNAME       string `bson:"username" json:"username"`
	BUYER_SKU_CODE string `bson:"buyer_sku_code" json:"buyer_sku_code"`
	COSTUMER_NO    string `bson:"customer_no" json:"customer_no"`
	REF_ID         string `bson:"ref_id" json:"ref_id"`
	SIGN           string `bson:"sign" json:"sign"`
}

type ResponseCheckStatus struct {
	BUYER_SKU_CODE string `bson:"buyer_sku_code" json:"buyer_sku_code"`
	COSTUMER_NO    string `bson:"customer_no" json:"customer_no"`
	MESSAGE        string `bson:"message" json:"message"`
	RC             string `bson:"rc" json:"rc"`
	REF_ID         string `bson:"ref_id" json:"ref_id"`
	SN             string `bson:"sn" json:"sn"`
	STATUS         string `bson:"status" json:"status"`
}

//End Check Status

// Start Inquire Pln

type InquirePLN struct {
	COMMANDS    string `bson:"commands" json:"commands"`
	COSTUMER_NO string `bson:"customer_no" json:"customer_no"`
}
type ResponseInquirePLN struct {
	COSTUMER_NO   string `bson:"customer_no" json:"customer_no"`
	METER_NO      string `bson:"meter_no" json:"meter_no"`
	SUBSCRIBER_ID string `bson:"subscriber_id" json:"subscriber_id"`
	NAME          string `bson:"name" json:"name"`
	SEGMENT_POWER string `bson:"segment_power" json:"segment_power"`
}

//End Inquire Pln
