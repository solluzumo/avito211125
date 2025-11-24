package models

type Filterable interface {
	FilterFieldMap() map[string]string
	TableName() string
}
