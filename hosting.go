package helpers

import (
	"bytes"

	"github.com/go-zoox/fetch"
)

func PlanHosting(url string, param, auth map[string]string) (*bytes.Buffer, error) {
	// api.version=1
	res, err := fetch.Get(url+"/json-api/matchpkgs",
		&fetch.Config{Headers: auth, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func CreateAccountHosting(url string, param, auth map[string]string) (*bytes.Buffer, error) {
	// api.version=1&username=username&domain=example.com&contactemail=example@gmail.com&pass=123456&plan=default
	res, err := fetch.Get(url+"/json-api/createacct",
		&fetch.Config{Headers: auth, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func UpdatePlanHosting(url string, param, auth map[string]string) (*bytes.Buffer, error) {
	// api.version=1&user=username&pkg=package1
	res, err := fetch.Get(url+"/json-api/changepackage",
		&fetch.Config{Headers: auth, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func DeleteHosting(url string, param, auth map[string]string) (*bytes.Buffer, error) {
	// api.version=1&username=username
	res, err := fetch.Get(url+"/json-api/removeacct",
		&fetch.Config{Headers: auth, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func SuspendHosting(url string, param, auth map[string]string) (*bytes.Buffer, error) {
	// api.version=1&user=username
	res, err := fetch.Get(url+"/json-api/suspendacct",
		&fetch.Config{Headers: auth, Params: param})
	return bytes.NewBuffer(res.Body), err
}

func UnsuspendHosting(url string, param, auth map[string]string) (*bytes.Buffer, error) {
	// api.version=1&user=username
	res, err := fetch.Get(url+"/json-api/unsuspendacct",
		&fetch.Config{Headers: auth, Params: param})
	return bytes.NewBuffer(res.Body), err
}
