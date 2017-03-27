package stack15

import (
	"fmt"
	"strings"
)

func mapContextToFields(context []interface{}, omitEmpty bool, fields map[string]string) {
	var v interface{}
	for i := 0; i < len(context); i += 2 {
		// handle odd key count
		if i == len(context)-1 {
			break
		}
		// skip nil values
		if v = context[i+1]; v == nil {
			continue
		}
		// expect string keys
		k, ok := context[i].(string)
		if !ok {
			// convert if not
			k = fmt.Sprintf("%v", context[i])
		}
		// caller is a special context key which is set when
		// the stack is inspected by a Caller* Handler
		if "caller" == k {
			ln := strings.Split(v.(string), ":")
			fields["file"] = ln[0]
			fields["line"] = ln[1]
			continue
		}
		// always use string value
		// (mixed types for the same field causes problems)
		var valueString string
		if s, ok := v.(string); ok {
			valueString = s
		} else {
			valueString = fmt.Sprintf("%v", v)
		}
		// optionally skip empty values
		if omitEmpty && len(valueString) == 0 {
			continue
		}
		fields[k] = valueString
	}
}
