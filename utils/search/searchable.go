package search

type Searchable interface {
	GetNext(userData interface{}) []Searchable
	GetRepr(userData interface{}) interface{}
}
