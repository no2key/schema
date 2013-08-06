package schema

import (
	"errors"
	"fmt"
)

var Builtin = struct {
	MaxLength ValidatorFunc
	Integer   ValidatorFunc
}{
	MaxLength: func(val interface{}, args []interface{}) error {
		stringVal, ok := val.(string)
		if !ok {
			return errors.New("Not a string")
		}

		maxLength := args[0].(int)
		if len(stringVal) > maxLength {
			return errors.New(fmt.Sprintf("String can't be longer than %v", maxLength))
		}
		return nil
	},
	Integer: func(val interface{}, args []interface{}) error {
		num, ok := val.(float64)
		if !ok {
			return errors.New("Not a number")
		}

		if float64(int(num)) != num {
			return errors.New(fmt.Sprintf("%v is not an integer", num))
		}
		return nil
	},
}
