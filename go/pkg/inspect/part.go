//ff:type feature=inspect type=model
//ff:what Part — one field of a record's words: its name, words, and whether it is head or an embedded record
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// Part is one field of a record's encoding, in word order. Head marks the TAG
// WORD and its extensions (Length, Issuer, format number); every other part is
// payload. Record marks an embedded record (a LOCAL Reference or a Value), whose
// first word is its own TAG WORD.
type Part struct {
	Name   string
	Words  []core.Word
	Head   bool
	Record bool
}
