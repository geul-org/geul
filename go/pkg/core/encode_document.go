//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what EncodeDocument — a whole document from HEAD and BODY records: DOC framing, BODY and END markers
package core

// EncodeDocument encodes head, the BODY marker, body and END and wraps them in
// the DOC record of the narrowest width (WrapDocument). The first record that
// does not encode is the error.
func EncodeDocument(head, body []Record) ([]byte, error) {
	in := []Word{}
	for _, r := range append(append(append([]Record{}, head...), Record{Kind: KindMeta, Codec: MetaBody}), body...) {
		ws, err := EncodeRecord(r)
		if err != nil {
			return nil, err
		}
		in = append(in, ws...)
	}
	return PutWords(WrapDocument(append(in, 0x180F))), nil
}
