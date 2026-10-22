//ff:func feature=core type=helper control=sequence
//ff:what TestCodecName — Meta, Value, ISSUER and unknown codecs
package core

import "testing"

func TestCodecName(t *testing.T) {
	if got := CodecName(KindMeta, MetaIDWidth); got != "ID-WIDTH" {
		t.Fatalf("meta: %q", got)
	}
	if got := CodecName(KindValue, CodecUTF8TextEven|1); got != "UTF8-TEXT" {
		t.Fatalf("value: %q", got)
	}
	if got := CodecName(KindEdge, CodecIssuer2Word); got != "ISSUER-2WORD" {
		t.Fatalf("issuer: %q", got)
	}
	if got := CodecName(KindEdge, 0x60); got != "" {
		t.Fatalf("profile edge: %q", got)
	}
	if got := CodecName(KindReference, 0x01); got != "" {
		t.Fatalf("reference: %q", got)
	}
}
