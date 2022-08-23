package core_es

type EsValueType struct {
	Type string
}

func Text() EsValueType {
	return EsValueType{Type: "text"}
}
func Keyword() EsValueType {
	return EsValueType{Type: "Keyword"}
}
func Float() EsValueType {
	return EsValueType{Type: "float"}
}
func Double() EsValueType {
	return EsValueType{Type: "double"}
}
func Long() EsValueType {
	return EsValueType{Type: "long"}
}
func Integer() EsValueType {
	return EsValueType{Type: "integer"}
}

func Nested() EsValueType {
	return EsValueType{Type: "nested"}
}
type esProperties struct {
	Field string
	Type EsValueType
}

type IEsMapping interface {
	SetType(field string , typeValue EsValueType) *esMapping
	SetCustomMapping(mapping map[string]interface{}) *esMapping
	SetNestedType(nestedField string ,field string , typeValue EsValueType) *esMapping
	Done() map[string]interface{}
}
type esMapping struct {
	Properties map[string]interface{}
}

func NewEsMapping() IEsMapping {
	return &esMapping{
		Properties: make(map[string]interface{}),
	}
}

func (es *esMapping) SetType(field string , typeValue EsValueType) *esMapping {
	if len(typeValue.Type) == 0 {
		return es
	}
	es.Properties[field] = map[string]interface{}{
		"type": typeValue.Type,
	}
	return es
}

func (es *esMapping) SetCustomMapping(mapping map[string]interface{}) *esMapping {
	es.Properties = mapping
	return es
}

func (es *esMapping) SetNestedType(nestedField string ,field string , typeValue EsValueType) *esMapping {
	if len(typeValue.Type) == 0 {
		return es
	}

	nestedObject,ok := es.Properties[nestedField].(map[string]interface{})
	if nestedObject != nil && ok {
		if nestedObject["properties"] == nil {
			nestedObject["properties"] = make(map[string]interface{})
		}
		properties, ok2 := nestedObject["properties"].(map[string]interface{})
		if ok2 {
			properties[field] =  map[string]string{
				"type": typeValue.Type,
			}
		}
	}
	return es
}

func (es *esMapping) Done() map[string]interface{} {
	return map[string]interface{}{
		"mappings": map[string]interface{}{
			"properties": es.Properties,
		},
	}
}