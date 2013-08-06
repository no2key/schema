package schema

import (
	"errors"
)

type ValidatorFunc func(val interface{}, args []interface{}) error

type Validator struct {
	required bool
	fun      ValidatorFunc
	args     []interface{}
}

func Validate(data, schema map[string]interface{}) error {
	return subValidate(data, schema, "")
}

func subValidate(data, schema map[string]interface{}, parentPath string) error {
	for key, value := range schema {
		if validator, ok := value.(Validator); ok {
			// Validate values
			if inputValue, ok := data[key]; ok {
				err := validator.fun(inputValue, validator.args)
				if err != nil {
					return errors.New(parentPath + key + ": " + err.Error())
				}
			} else if validator.required {
				return errors.New(parentPath + key + ": value is required")
			}
		} else {
			// Not a validator, then it must be a map[string]interface{}
			subSchema, ok := value.(map[string]interface{})
			if !ok {
				return errors.New(parentPath + "." + key + ": must be either a Validator or map[string]interface{}")
			}
			if subData, ok := data[key].(map[string]interface{}); ok {
				err := subValidate(subData, subSchema, parentPath+key+".")
				if err != nil {
					return err
				}
			} else {
				return errors.New("value is required")
			}
		}
	}
	return nil
}
