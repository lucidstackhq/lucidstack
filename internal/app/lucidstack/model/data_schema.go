package model

import (
	"fmt"
	"regexp"
)

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

func (c *Constraints) ValidateString(val string) error {
	if c.Pattern != nil {
		reg, err := regexp.Compile(*c.Pattern)

		if err != nil {
			return err
		}

		if !reg.MatchString(val) {
			return fmt.Errorf("string value %s does not match pattern %s", val, *c.Pattern)
		}
	}

	if c.MaxLength != nil {
		if len(val) > int(*c.MaxLength) {
			return fmt.Errorf("string value %s is too long", val)
		}
	}

	if c.MinLength != nil {
		if *c.MinLength < *c.MaxLength {
			return fmt.Errorf("string value %s is too short", val)
		}
	}

	return nil
}

func (c *Constraints) ValidateInteger(val int64) error {
	if c.Max != nil {
		maxInt, ok := c.Max.(int64)

		if !ok {
			return fmt.Errorf("max value is not an integer")
		}

		if val > maxInt {
			return fmt.Errorf("integer value %d is greater than max value %d", val, maxInt)
		}
	}

	if c.Min != nil {
		minInt, ok := c.Min.(int64)

		if !ok {
			return fmt.Errorf("min value is not an integer")
		}

		if val < minInt {
			return fmt.Errorf("integer value %d is less than min value %d", val, minInt)
		}
	}

	return nil
}

func (c *Constraints) ValidateFloat(val float64) error {
	if c.Max != nil {
		maxFloat, ok := c.Max.(float64)

		if !ok {
			return fmt.Errorf("max value is not a float")
		}

		if val > maxFloat {
			return fmt.Errorf("float value %f is greater than max value %f", val, maxFloat)
		}
	}

	if c.Min != nil {
		minFloat, ok := c.Min.(float64)

		if !ok {
			return fmt.Errorf("min value is not a float")
		}

		if val < minFloat {
			return fmt.Errorf("float value %f is less than min value %f", val, minFloat)
		}
	}

	return nil
}

func (c *Constraints) ValidateObject(val map[string]interface{}) error {
	if c.MaxLength != nil {
		if len(val) > int(*c.MaxLength) {
			return fmt.Errorf("object value %s is too long", val)
		}
	}

	if c.MinLength != nil {
		if *c.MinLength < *c.MaxLength {
			return fmt.Errorf("object value %s is too short", val)
		}
	}

	return nil
}

func (c *Constraints) ValidateArray(val []interface{}) error {
	if c.MaxLength != nil {
		if len(val) > int(*c.MaxLength) {
			return fmt.Errorf("array value %s is too long", val)
		}
	}

	if c.MinLength != nil {
		if *c.MinLength < *c.MaxLength {
			return fmt.Errorf("array value %s is too short", val)
		}
	}

	return nil
}

func (d *DataSchema) Validate(data interface{}) error {
	if d == nil {
		return nil
	}

	switch d.Type {
	case DataTypeString:
		v, ok := data.(string)

		if !ok {
			return fmt.Errorf("data is not a string")
		}

		if d.Constraints != nil {
			if err := d.Constraints.ValidateString(v); err != nil {
				return err
			}
		}

		break

	case DataTypeInteger:
		v, ok := data.(int64)

		if !ok {
			return fmt.Errorf("data is not an integer")
		}

		if d.Constraints != nil {
			if err := d.Constraints.ValidateInteger(v); err != nil {
				return err
			}
		}

		break

	case DataTypeFloat:
		v, ok := data.(float64)

		if !ok {
			return fmt.Errorf("data is not a float")
		}

		if _, ok := data.(float64); !ok {
			return fmt.Errorf("data is not a float")
		}

		if d.Constraints != nil {
			if err := d.Constraints.ValidateFloat(v); err != nil {
				return err
			}
		}

		break

	case DataTypeBool:
		if _, ok := data.(bool); !ok {
			return fmt.Errorf("data is not a boolean")
		}

		break

	case DataTypeObject:
		v, ok := data.(map[string]interface{})

		if !ok {
			return fmt.Errorf("data is not a object")
		}

		for key, value := range v {
			fieldSchema, ok := d.Fields[key]
			if !ok {
				return fmt.Errorf("schema for field %s not found", key)
			}

			err := fieldSchema.Validate(value)
			if err != nil {
				return err
			}
		}

		if d.Constraints != nil {
			if err := d.Constraints.ValidateObject(v); err != nil {
				return err
			}
		}

		break

	case DataTypeArray:
		v, ok := data.([]interface{})
		if !ok {
			return fmt.Errorf("data is not an array")
		}

		for _, value := range v {
			err := d.Elements.Validate(value)
			if err != nil {
				return err
			}
		}

		if d.Constraints != nil {
			if err := d.Constraints.ValidateArray(v); err != nil {
				return err
			}
		}

		break

	default:
		return fmt.Errorf("unknown data type %s", d.Type)
	}

	return nil
}
