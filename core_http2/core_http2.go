package core_http2

import (
	"encoding/json"
	"fmt"
	"github.com/smart-think-app/backend-core/core_error"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
	"strings"
)

type proxy struct {
	domain string
	client *http.Client
}

func NewProxy(domain string) *proxy {
	client := &http.Client{}
	client.Transport = &http2.Transport{}
	return &proxy{
		domain: domain,
		client: client,
	}
}
func(p *proxy) GetMethod(endpoint string, response interface{}) error{
	req , err := http.NewRequest("GET",fmt.Sprintf("%s/%s" ,
		strings.TrimRight(p.domain,"/") ,strings.TrimLeft(endpoint,"/")),nil)
	resp, err := p.client.Do(req)
	if err != nil {
		err = core_error.NewCoreError().InternalError(err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = core_error.NewCoreError().InternalError(err.Error())
		return err
	}
	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(body , response)
		if err != nil {
			err = core_error.NewCoreError().InternalError(err.Error())
			return err
		}
		return nil
	}
	err = core_error.NewCoreError().CustomError(resp.Status , resp.StatusCode)
	return err
}