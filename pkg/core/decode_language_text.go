//ff:func feature=core type=decoder control=sequence
//ff:what DecodeLanguageText — the language record and text bytes of a UTF8-LANGUAGE-TEXT Value
package core

import "errors"

// DecodeLanguageText splits a UTF8-LANGUAGE-TEXT Value r into its language
// record (the first record of the Payload) and its text bytes (the rest, the
// odd member's final padding byte dropped). It fails for another codec, a
// Payload that does not start with a whole record, or a bad odd padding.
func DecodeLanguageText(r Record) (Record, []byte, error) {
	if r.Kind != KindValue || r.Codec&^1 != CodecUTF8LanguageTextEven {
		return Record{}, nil, errors.New("geul/core: not a UTF8-LANGUAGE-TEXT Value")
	}
	lang, n, err := DecodeRecord(r.Payload)
	if err != nil {
		return Record{}, nil, err
	}
	text, err := DecodeBytesValue(Record{Kind: KindValue, Codec: r.Codec&1 | CodecUTF8TextEven, Payload: r.Payload[n:]}, CodecUTF8TextEven)
	if err != nil {
		return Record{}, nil, err
	}
	return lang, text, nil
}
