package core_es

import (
	"github.com/smart-think-app/backend-core/core_http2"
)

type IEsCore interface {
	CreateIndex(index string) error
	CreateIndexWithMapping(index string, mapping interface{}) error
	SetAuthorization(token string)
}
type esCore struct {
	EsProxy core_http2.ICoreHTTP2
	Config struct {
		Authorization string
	}
}

func NewEsCore(domainEs string) IEsCore {
	return &esCore{
		EsProxy: core_http2.NewProxy(domainEs),
	}
}

func (es *esCore) SetAuthorization(token string) {
	es.Config.Authorization = token
	es.EsProxy.AddHeader("Authorization" , es.Config.Authorization)
}

func(es *esCore) CreateIndex(index string) error {
	err := es.EsProxy.PutMethod(index , nil,nil,nil,[]string{"Authorization"})
	if err != nil {
		return err
	}
	return nil
}

func(es *esCore) CreateIndexWithMapping(index string, mapping interface{}) error {
	return nil
}