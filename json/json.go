package json

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/pretty"
)

var (
	JJ     = jsoniter.ConfigCompatibleWithStandardLibrary // 100% compatible with standard library behavior
	FastJJ = jsoniter.ConfigFastest                       // ConfigFastest marshals float with only 6 digits precision
)

func Pretty(msg string) string {
	return fmt.Sprintf("%s", pretty.Pretty([]byte(msg)))
}
