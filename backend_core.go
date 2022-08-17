package backend_core

type ICoreHTTP2 interface {
	GetMethod(endpoint string, response interface{}) error
}