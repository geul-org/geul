//ff:func feature=resolve type=decoder control=sequence
//ff:what metaTime — the instant a time Meta (EFFECTIVE-FROM/UNTIL) carries in its embedded Value
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// metaTime reads the Value embedded in a time Meta record. A payload that is not
// one GREGORIAN-DATE or POSIX-TIMESTAMP-SECONDS Value fails verification.
func metaTime(r core.Record) (timeValue, error) {
	v, n, err := core.DecodeRecord(r.Payload)
	if err != nil || n != len(r.Payload) {
		return timeValue{}, fmt.Errorf("%w: %s payload is not one Value", ErrVerify, core.CodecName(r.Kind, r.Codec))
	}
	tm, err := core.DecodeTimeValue(v)
	if err != nil {
		return timeValue{}, fmt.Errorf("%w: %v", ErrVerify, err)
	}
	return timeValue{tm}, nil
}
