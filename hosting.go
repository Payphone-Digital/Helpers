package helpers

import (
	"bytes"

	"github.com/go-zoox/fetch"
)

func Hosting(url string, header map[string]string) *HostingNeed {
	return &HostingNeed{
		Url:    url,
		Header: header,
	}
}

func (h *HostingNeed) Plan(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(h.Url+"/json-api/matchpkgs",
		&fetch.Config{Headers: h.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (h *HostingNeed) CreateAccount(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(h.Url+"/json-api/createacct",
		&fetch.Config{Headers: h.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (h *HostingNeed) UpdatePlan(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(h.Url+"/json-api/changepackage",
		&fetch.Config{Headers: h.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (h *HostingNeed) Delete(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(h.Url+"/json-api/removeacct",
		&fetch.Config{Headers: h.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (h *HostingNeed) Suspend(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(h.Url+"/json-api/suspendacct",
		&fetch.Config{Headers: h.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func (h *HostingNeed) Unsuspend(param map[string]string) (*bytes.Buffer, error) {
	res, err := fetch.Get(h.Url+"/json-api/unsuspendacct",
		&fetch.Config{Headers: h.Header, Params: param})
	return bytes.NewBuffer(res.Body), err
}
