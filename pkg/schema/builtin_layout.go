//ff:func feature=schema type=rule control=selection
//ff:what builtinLayout — E4: decode a core-embedded Edge (version list records, codec spec vocabulary); the spec it must follow and the error
package schema

import "github.com/geul-org/gogeul/pkg/core"

// builtinLayout decodes r when it is a core-embedded Edge (GEUL-ORG 0층): the
// version list records and DECLARE, FIELD, ENUM and CODEC. It returns the spec
// document r must follow and the decode error; nil for a fitting or another record.
func builtinLayout(r core.Record) (string, error) {
	var err error
	switch r.Codec {
	case core.EdgePublication, core.EdgeDeprecation, core.EdgeTombstone:
		_, err = core.DecodePublication(r)
		return "edge/PUBLICATION 레이아웃", err
	case core.EdgeDeclare:
		_, err = core.DecodeDeclare(r)
		return "edge/DECLARE 레이아웃", err
	case core.EdgeField:
		_, err = core.DecodeFieldDecl(r)
		return "edge/FIELD 레이아웃", err
	case core.EdgeEnum:
		_, err = core.DecodeEnumDecl(r)
		return "edge/ENUM 레이아웃", err
	case core.EdgeCodec:
		_, err = core.DecodeCodecDecl(r)
		return "edge/CODEC 레이아웃", err
	}
	return "", nil
}
