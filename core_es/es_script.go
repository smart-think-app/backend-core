package core_es

type esScript struct {
	Source []string
	Size int
	From int
	Query map[string]interface{}
}

type IEsScript interface {
	SetSize(size int) *esScript
	SetSource(source []string) *esScript
	SetFrom(from int) *esScript
	SetCustomQuery(query map[string]interface{}) *esScript
	Done() map[string]interface{}
}

func NewEsScript() IEsScript {
	query := make(map[string]interface{})
	query["query"] = make(map[string]interface{})
	return &esScript{
		Source: nil,
		Size:   0,
		From:   0,
		Query:  query,
	}
}

func (es *esScript) SetSize(size int) *esScript {
	es.Size = size
	return es
}

func (es *esScript) SetSource(source []string) *esScript {
	es.Source = source
	return es
}

func (es *esScript) SetFrom(from int) *esScript {
	es.From = from
	return es
}

func (es *esScript) SetCustomQuery(query map[string]interface{}) *esScript {
	es.Query = query
	return es
}

func (es *esScript) Done() map[string]interface{}{
	script := make(map[string]interface{})

	script["size"] = es.Size
	script["_source"] = es.Source
	script["from"] = es.From
	script["query"] = es.Query
	return script
}