package model

type DataSchema struct {
	Type        DataType               `json:"type"`
	Fields      map[string]*DataSchema `json:"fields"`
	Elements    *DataSchema            `json:"elements"`
	Constraints *Constraints           `json:"constraints"`
}

type DataType string

const (
	DataTypeString  DataType = "string"
	DataTypeInteger DataType = "integer"
	DataTypeFloat   DataType = "float"
	DataTypeBool    DataType = "bool"
	DataTypeObject  DataType = "object"
	DataTypeArray   DataType = "array"
)

type Constraints struct {
	Min       interface{} `json:"min"`
	Max       interface{} `json:"max"`
	MinLength *int64      `json:"min_length"`
	MaxLength *int64      `json:"max_length"`
	Pattern   *string     `json:"pattern"`
}
