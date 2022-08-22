package core_http2

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/smart-think-app/backend-core/core_error"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type ICoreHTTP2 interface {
	GetMethod(endpoint string, response interface{},customHeader []CoreHeader,baseHeader []string ) error
	PostMethod(endpoint string, response interface{},payload interface{},customHeader []CoreHeader,baseHeader []string ) error
	PutMethod(endpoint string, response interface{},payload interface{},customHeader []CoreHeader,baseHeader []string ) error
	AddHeader(header string , value string) *proxy
}

type proxy struct {
	domain string
	client *http.Client
	headerMap map[string][]string
}
type CoreHeader struct {
	HeaderKey string
	HeaderValue []string
}

func NewProxy(domain string) ICoreHTTP2 {
	client := &http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &proxy{
		domain: domain,
		client: client,
		headerMap: make(map[string][]string),
	}
}
func(p *proxy) AddHeader(header string , value string) *proxy{
	headerList := p.headerMap[header]
	headerList = append(headerList , value)
	p.headerMap[header] = headerList
	return p
}
func (p *proxy)setHeader(customHeader []CoreHeader,baseHeader []string) http.Header {
	header := http.Header{
		"Content-Type": {"application/json"},
	}
	if baseHeader != nil {
		for _, value := range baseHeader {
			if len(value) != 0 && p.headerMap[value] != nil {
				header[value] =  p.headerMap[value]
			}
		}
	}
	if customHeader != nil  {
		for _, value := range customHeader {
			if len(value.HeaderKey) != 0 {
				header[value.HeaderKey] = value.HeaderValue
			}
		}
	}
	return header
}
func(p *proxy) GetMethod(endpoint string, response interface{},customHeader []CoreHeader,baseHeader []string ) error{
	req , err := http.NewRequest("GET",fmt.Sprintf("%s/%s" ,
		strings.TrimRight(p.domain,"/") ,strings.TrimLeft(endpoint,"/")),nil)
	if err != nil {
		err = core_error.NewCoreError().InternalError(err.Error())
		return err
	}
	req.Header = p.setHeader(customHeader , baseHeader)

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

func(p *proxy) PostMethod(endpoint string, response interface{},payload interface{},customHeader []CoreHeader,baseHeader []string ) error{
	jsonData, err := json.Marshal(payload)
	if err != nil {
		err = core_error.NewCoreError().InternalError(err.Error())
		return err
	}

	req , err := http.NewRequest("POST",fmt.Sprintf("%s/%s" ,
		strings.TrimRight(p.domain,"/") ,strings.TrimLeft(endpoint,"/")),
		bytes.NewBuffer(jsonData))
	if err != nil {
		err = core_error.NewCoreError().InternalError(err.Error())
		return err
	}
	req.Header = p.setHeader(customHeader , baseHeader)

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

func(p *proxy) PutMethod(endpoint string, response interface{},payload interface{},customHeader []CoreHeader,baseHeader []string ) error{
	var bodyRequest io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			err = core_error.NewCoreError().InternalError(err.Error())
			return err
		}
		bodyRequest = bytes.NewBuffer(jsonData)
	}

	req , err := http.NewRequest("PUT",fmt.Sprintf("%s/%s" ,
		strings.TrimRight(p.domain,"/") ,strings.TrimLeft(endpoint,"/")),
		bodyRequest)
	if err != nil {
		err = core_error.NewCoreError().InternalError(err.Error())
		return err
	}
	req.Header = p.setHeader(customHeader , baseHeader)

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
		if response != nil {
			err = json.Unmarshal(body , response)
			if err != nil {
				err = core_error.NewCoreError().InternalError(err.Error())
				return err
			}
		}
		return nil
	}
	err = core_error.NewCoreError().CustomError(resp.Status , resp.StatusCode)
	return err
}