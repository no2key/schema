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

/*func main() {

	// Validators
	var Max10 ValidatorFunc

	Max10 = func(val interface{}, args []interface{}) error {
		f, ok := val.(float64)
		if !ok {
			return errors.New("Not a number")
		}
		if f > 10 {
			return errors.New(fmt.Sprintf("Input value can't be more than 10 (was %v)", f))
		}
		return nil
	}

	var MaxNum ValidatorFunc

	MaxNum = func(val interface{}, args []interface{}) error {
		f, ok := val.(float64)
		if !ok {
			return errors.New("Not a number")
		}

		max := args[0].(float64)
		if f > max {
			return errors.New(fmt.Sprintf("Input value can't be more than %v (was %v)", max, f))
		}
		return nil
	}

	// Dummy data to validate
	data := []byte(`{
		"number1": 8,
		"number2": 10,
		"hugeNumber": 1000,
		"extraNumbers": {
		}
	}`)

	schema := map[string]interface{}{
		"number1":    Validator{required: true, fun: Max10},
		"number2":    Validator{required: false, fun: Max10},
		"hugeNumber": Validator{required: true, fun: MaxNum, args: []interface{}{5000.0}},
		"extraNumbers": map[string]interface{}{
			"one": Validator{required: false, fun: Max10},
			"two": Validator{required: true, fun: Max10},
		},
	}

	var v map[string]interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	err = Validate(v, schema)
	if err != nil {
		log.Fatal(err)
	}
}
*/
