//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what encodeVersionList — a version list document: SPEC-VERSION, PARENT, RECORDED-AT, then the records
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// encodeVersionList writes a version list document (edge/PUBLICATION): HEAD
// SPEC-VERSION, PARENT (the previous list's content hash, when there is one)
// and RECORDED-AT (seconds), then the BODY records in order.
func encodeVersionList(records []core.Record, parent *[32]byte, at int64) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	in := enc(core.EncodeSpecVersion(core.SpecVersion))
	if parent != nil {
		h, _ := core.ReadWords(parent[:])
		in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: 0x30, Payload: enc(core.Record{Kind: core.KindValue, Codec: core.CodecSHA256Digest, Payload: h})})...)
	}
	in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: 0x25, Payload: enc(core.EncodePosixSeconds(at))})...)
	in = append(in, 0x180E)
	for _, r := range records {
		in = append(in, enc(r)...)
	}
	return core.PutWords(core.WrapDocument(append(in, 0x180F)))
}
