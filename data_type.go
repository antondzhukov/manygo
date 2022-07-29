package manygo

type DataType string

const (
	DataTypeUndefined DataType = ""
	DataTypeText      DataType = "text"
	DataTypeNumber    DataType = "number"
	DataTypeDate      DataType = "date"
	DataTypeDateTime  DataType = "datetime"
	DataTypeBoolean   DataType = "boolean"
)
