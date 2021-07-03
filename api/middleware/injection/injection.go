package injection

import (
	"encoding/json"

	"github.com/fanfit/user-service/api/authentication"
)

// Reset takes all injected variables to their original values
func Reset() {
	Unmarshal = unmarshalDefault
	VerifyAccessToken = verifyAccessTokenDefault
}

// Unmarshal decodes a byte stream into the target interface
var Unmarshal = unmarshalDefault
var unmarshalDefault = json.Unmarshal

var VerifyAccessToken = verifyAccessTokenDefault
var verifyAccessTokenDefault = authentication.CheckJwt
