//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterEncodeRecord — Stored keeps raw; a smaller frame is zstd; a larger one stays raw; codec errors fail
package pack

import (
	"errors"
	"testing"
)

func TestWriterEncodeRecord(t *testing.T) {
	w := &writer{o: BuildOptions{Codec: Stored{}}}
	if enc, b, err := w.encodeRecord([]byte("abc")); err != nil || enc != encRaw || string(b) != "abc" {
		t.Fatal(enc, err)
	}
	w.o.Codec = fakeCodec{out: []byte{1}}
	if enc, b, _ := w.encodeRecord([]byte("abc")); enc != encZstd || len(b) != 1 {
		t.Fatal(enc)
	}
	w.o.Codec = fakeCodec{out: []byte("abcd")}
	if enc, _, _ := w.encodeRecord([]byte("abc")); enc != encRaw {
		t.Fatal(enc)
	}
	w.o.Codec = fakeCodec{err: errors.New("x")}
	if _, _, err := w.encodeRecord([]byte("abc")); err == nil {
		t.Fatal("err")
	}
}
