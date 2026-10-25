//ff:func feature=core type=helper control=sequence
//ff:what TestCodecName — Reference, Meta, Value, execution Edge, ISSUER and unknown codecs
package core

import "testing"

func TestCodecName(t *testing.T) {
	if got := CodecName(KindMeta, MetaImportDoc); got != "IMPORT-DOC" {
		t.Fatalf("meta: %q", got)
	}
	if got := CodecName(KindMeta, MetaEffectiveFrom); got != "EFFECTIVE-FROM" {
		t.Fatalf("effective-from: %q", got)
	}
	if got := CodecName(KindMeta, 0x15); got != "" {
		t.Fatalf("reserved: %q", got)
	}
	if got := CodecName(KindValue, CodecUTF8TextEven|1); got != "UTF8-TEXT" {
		t.Fatalf("value: %q", got)
	}
	if got := CodecName(KindValue, CodecPosixSeconds); got != "POSIX-TIMESTAMP-SECONDS" {
		t.Fatalf("time value: %q", got)
	}
	if got := CodecName(KindValue, CodecEDTFEven|1); got != "EDTF" {
		t.Fatalf("EDTF odd: %q", got)
	}
	if got := CodecName(KindValue, 0x26); got != "" {
		t.Fatalf("reserved value: %q", got)
	}
	if got := CodecName(KindEdge, CodecIssuer2Word); got != "ISSUER-2WORD" {
		t.Fatalf("issuer: %q", got)
	}
	if got := CodecName(KindEdge, 0x41); got != "FUNCTION-DECLARE" {
		t.Fatalf("execution edge: %q", got)
	}
	if got := CodecName(KindEdge, EdgeBind); got != "BIND" {
		t.Fatalf("bind: %q", got)
	}
	if got := CodecName(KindEdge, 0x40); got != "" {
		t.Fatalf("removed SYMBOL-DECLARE: %q", got)
	}
	if got := CodecName(KindEdge, 0x60); got != "" {
		t.Fatalf("profile edge: %q", got)
	}
	if got := CodecName(KindReference, RefLocal<<3|4); got != "LOCAL" {
		t.Fatalf("reference: %q", got)
	}
	if got := CodecName(KindReference, 0x07); got != "" {
		t.Fatalf("reserved reference: %q", got)
	}
}
