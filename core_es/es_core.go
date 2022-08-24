package core_es

import (
	"fmt"
	"github.com/smart-think-app/backend-core/core_http2"
)

type IEsCore interface {
	CreateIndex(index string) error
	CreateIndexWithMapping(index string, mapping IEsMapping) error
	SetAuthorization(token string)
	CreateDocument(index string, document interface{}) error
	SearchDocument(index string, query interface{}) (*EsResponse,error)
	SearchDocumentAdvance(index string, script IEsScript) (*EsResponse,error)
}
type esCore struct {
	EsProxy core_http2.ICoreHTTP2
	Config  struct {
		Authorization string
	}
}

type EsResponse struct {
	Took     float64 `json:"took"`
	TimedOut bool    `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string      `json:"_index"`
			Id     string      `json:"_id"`
			Score  float64     `json:"_score"`
			Source map[string]interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func NewEsCore(domainEs string) IEsCore {
	return &esCore{
		EsProxy: core_http2.NewProxy(domainEs),
	}
}

func (es *esCore) SetAuthorization(token string) {
	es.Config.Authorization = token
	es.EsProxy.AddHeader("Authorization", es.Config.Authorization)
}

func (es *esCore) CreateIndex(index string) error {
	err := es.EsProxy.PutMethod(index, nil, nil, nil, []string{"Authorization"})
	if err != nil {
		return err
	}
	return nil
}

func (es *esCore) CreateIndexWithMapping(index string, mapping IEsMapping) error {
	err := es.EsProxy.PutMethod(index, nil, mapping.Done(), nil, []string{"Authorization"})
	if err != nil {
		return err
	}
	return nil
}

func (es *esCore) CreateDocument(index string, document interface{}) error {
	endpoint := fmt.Sprintf("%s/_doc", index)
	err := es.EsProxy.PostMethod(endpoint, nil, document, nil, []string{"Authorization"})
	if err != nil {
		return err
	}
	return nil
}

func (es *esCore) SearchDocument(index string, query interface{}) (*EsResponse,error) {

	if query == nil {
		query = make(map[string]interface{})
	}
	endPoint := fmt.Sprintf("%s/_search" , index)
	var response EsResponse
	err := es.EsProxy.PostMethod(endPoint, &response , query,nil,[]string{"Authorization"})
	if err != nil {
		return nil ,err
	}

	return &response, nil
}

func (es *esCore) SearchDocumentAdvance(index string, script IEsScript) (*EsResponse,error) {


	endPoint := fmt.Sprintf("%s/_search" , index)
	var response EsResponse
	err := es.EsProxy.PostMethod(endPoint, &response , script.Done(),nil,[]string{"Authorization"})
	if err != nil {
		return nil ,err
	}

	return &response, nil
}
