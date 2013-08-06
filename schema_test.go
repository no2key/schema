package schema

import (
	"encoding/json"
	"testing"
)

func TestBasic(t *testing.T) {
	data := []byte(`{
		"title": "A basic test",
		"intValue": 50
    }`)

	var v map[string]interface{}
	err := json.Unmarshal(data, &v)

	schema := map[string]interface{}{
		"title":    Validator{required: true, fun: Builtin.MaxLength, args: []interface{}{25}},
		"intValue": Validator{required: true, fun: Builtin.Integer},
	}

	err = Validate(v, schema)
	if err != nil {
		panic(err)
	}
}
