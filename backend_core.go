package backend_core

import "github.com/smart-think-app/backend-core/core_http2"

type ICoreHTTP2 interface {
	GetMethod(endpoint string, response interface{},customHeader []core_http2.CoreHeader,baseHeader []string ) error
	PostMethod(endpoint string, response interface{},payload interface{},customHeader []core_http2.CoreHeader,baseHeader []string ) error
}