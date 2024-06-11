package xpretty

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

// PrettyJsonArray
// Deprecated: use PrettyJson instead.
// func PrettyJsonArray(rawJson string, opts ...PrettyOptFunc) error {
// 	opt := PrettyOpts{indent: 2}
// 	bindPrettyOpts(&opt, opts...)

// 	var obj []interface{}
// 	e := json.Unmarshal([]byte(rawJson), &obj)

// 	f := colorjson.NewFormatter()
// 	f.Indent = opt.indent

// 	s, _ := f.Marshal(obj)
// 	fmt.Println(string(s))
// 	return e
// }

func PrettyJSON(rawJson string, opts ...PrettyOptFunc) error {
	return pretty([]byte(rawJson), opts...)
}

func pretty(raw []byte, opts ...PrettyOptFunc) error {
	opt := PrettyOpts{indent: 2}
	bindPrettyOpts(&opt, opts...)

	obj, err := getObj(raw)
	if err != nil {
		return err
	}

	f := colorjson.NewFormatter()
	f.Indent = opt.indent

	s, err := f.Marshal(obj)
	fmt.Println(string(s))
	return err
}

func getObj(raw []byte) (interface{}, error) {
	var obj map[string]interface{}

	err := json.Unmarshal([]byte(raw), &obj)
	if err != nil {
		var objArr []interface{}
		err = json.Unmarshal([]byte(raw), &objArr)
		if err != nil {
			return nil, err
		}

		return objArr, nil
	}

	return obj, nil
}

// PrettyStruct print struct as json str.
func PrettyStruct(data any) error {
	bt, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return pretty(bt)
}

func PrettyMap(obj map[string]interface{}, opts ...PrettyOptFunc) error {
	opt := PrettyOpts{indent: 2}
	bindPrettyOpts(&opt, opts...)

	f := colorjson.NewFormatter()
	f.Indent = opt.indent

	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
	return nil
}
