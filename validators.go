package schema

import (
	"errors"
	"fmt"
	"time"
)

var Builtin = struct {
	MaxLength ValidatorFunc
	Integer   ValidatorFunc
	Time      ValidatorFunc
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
	Time: func(val interface{}, args []interface{}) error {
		if len(args) != 1 {
			return errors.New("Time only takes one argument (layout)")
		}

		timeValue, ok := val.(string)
		if !ok {
			return errors.New("Time value must be a string")
		}

		layout, ok := args[0].(string)
		if !ok {
			return errors.New("Layout must be a string")
		}

		_, err := time.Parse(layout, timeValue)
		if err != nil {
			return err
		}

		return nil
	},
}
