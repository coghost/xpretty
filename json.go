package xpretty

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

func PrettyJson(rawJson string, opts ...PrettyOptFunc) {
	opt := PrettyOpts{indent: 2}
	bindPrettyOpts(&opt, opts...)

	var obj map[string]interface{}
	json.Unmarshal([]byte(rawJson), &obj)

	f := colorjson.NewFormatter()
	f.Indent = opt.indent

	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}
