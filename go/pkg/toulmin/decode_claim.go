//ff:func feature=toulmin type=decoder control=sequence
//ff:what decodeClaim — read a Claim block at *pos given its object kind from the NODE header
package toulmin

import "github.com/geul-org/geul/go/pkg/core"

// decodeClaim reads a Claim block at ws[*pos] (subject REF, predicate word,
// object of kind, mode word) and advances *pos. kind comes from the NODE header.
func decodeClaim(ws []core.Word, pos *int, kind ObjKind, w Wire) (Claim, error) {
	subject, err := TakeID(ws, pos, w.ID)
	if err != nil {
		return Claim{}, err
	}
	pred, err := TakeWord(ws, pos, "Claim predicate")
	if err != nil {
		return Claim{}, err
	}
	c := Claim{Subject: subject, Predicate: uint16(pred), Kind: kind}
	if err := decodeClaimObject(ws, pos, kind, w, &c); err != nil {
		return Claim{}, err
	}
	mode, err := TakeWord(ws, pos, "Claim mode")
	if err != nil {
		return Claim{}, err
	}
	c.Mode = uint16(mode)
	return c, nil
}
