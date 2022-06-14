package lib

import "encoding/json"

// Merge a struct to another struct
func Merge(from interface{}, to interface{}) error {
	j, e := json.Marshal(from)
	if nil == e {
		e = json.Unmarshal(j, to)
	}

	return e
}
