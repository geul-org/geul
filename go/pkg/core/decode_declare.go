//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what DecodeDeclare — a DECLARE Edge as Declare
package core

import "fmt"

// DecodeDeclare reads a DECLARE Edge (GEUL t002 docs/edge/DECLARE): the LOCAL
// Edge Id, the target word, the Issuer and format number of an ISSUER target,
// then LOCAL References to FIELD Edges up to the end.
func DecodeDeclare(r Record) (Declare, error) {
	if r.Kind != KindEdge || r.Codec != EdgeDeclare {
		return Declare{}, fmt.Errorf("geul/core: not a DECLARE Edge (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	id, pos, err := takeLocal(r.Payload)
	if err != nil {
		return Declare{}, err
	}
	if pos >= len(r.Payload) {
		return Declare{}, fmt.Errorf("geul/core: DECLARE without a target")
	}
	d := Declare{EdgeID: id, Target: r.Payload[pos]}
	pos++
	if w := IssuerWidth(SplitTarget(d.Target)); w > 0 {
		if len(r.Payload) < pos+w+1 {
			return Declare{}, fmt.Errorf("geul/core: DECLARE ISSUER target needs %d Issuer words and a format number", w)
		}
		d.Issuer, d.Format = r.Payload[pos:pos+w], r.Payload[pos+w]
		pos += w + 1
	}
	for pos < len(r.Payload) {
		f, n, err := takeLocal(r.Payload[pos:])
		if err != nil {
			return Declare{}, err
		}
		d.Fields = append(d.Fields, f)
		pos += n
	}
	return d, nil
}
