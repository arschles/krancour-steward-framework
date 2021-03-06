package lib

import (
	"fmt"
	"strings"
)

var (
	emptyJSONObject = JSONObject(map[string]interface{}{})
)

type errMalformedKV struct {
	kv []string
}

func (e errMalformedKV) Error() string {
	return fmt.Sprintf("malformed JSONObject-encoded key/value pair %s", e.kv)
}

// JSONObject is a convenience wrapper around a Go type that represents a JSON object
type JSONObject map[string]interface{}

// EmptyJSONObject returns an empty JSONObject
func EmptyJSONObject() JSONObject {
	return emptyJSONObject
}

// MarshalText is the encoding.TextMarshaler implementation
func (j JSONObject) EncodeToString() string {
	slc := make([]string, len(j))
	i := 0
	for key, val := range j {
		slc[i] = fmt.Sprintf("%s=%s", key, val)
		i++
	}
	return strings.Join(slc, ",")
}

// JSONObjectFromString decodes a string into a JSONObject. Returns a non-nil error if the string was not a valid JSONObject
func JSONObjectFromString(str string) (JSONObject, error) {
	if len(str) == 0 {
		return JSONObject(map[string]interface{}{}), nil
	}
	mp := map[string]interface{}{}
	spl := strings.Split(str, ",")
	for _, s := range spl {
		kv := strings.Split(s, "=")
		if len(kv) != 2 {
			return JSONObject(map[string]interface{}{}), errMalformedKV{kv: kv}
		}
		key, val := kv[0], kv[1]
		mp[key] = val
	}
	return JSONObject(mp), nil
}
