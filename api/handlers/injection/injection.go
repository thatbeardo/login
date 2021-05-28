package injection

import (
	"encoding/json"
)

// Reset takes all injected variables to their original values
func Reset() {
	Marshal = marshalDefault
	Unmarshal = unmarshalDefault
}

// Marshal returns the JSON encoding of the parameter passed
var Marshal = marshalDefault
var marshalDefault = json.Marshal

// Unmarshal encodes the byte stream into json
var Unmarshal = unmarshalDefault
var unmarshalDefault = json.Unmarshal
