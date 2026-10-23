//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeTimeValue — dates and timestamps including negatives; wrong codec and width fail
package core

import (
	"testing"
	"time"
)

func TestDecodeTimeValue(t *testing.T) {
	day := func(ws ...Word) Record { return Record{Kind: KindValue, Codec: CodecGregorianDate, Payload: ws} }
	if tm, err := DecodeTimeValue(day(0, 1)); err != nil || !tm.Equal(time.Date(1970, 1, 2, 0, 0, 0, 0, time.UTC)) {
		t.Fatalf("date: %v %v", tm, err)
	}
	if tm, err := DecodeTimeValue(day(0xffff, 0xffff)); err != nil || !tm.Equal(time.Date(1969, 12, 31, 0, 0, 0, 0, time.UTC)) {
		t.Fatalf("negative date: %v %v", tm, err)
	}
	sec := Record{Kind: KindValue, Codec: CodecPosixSeconds, Payload: []Word{0xffff, 0xffff, 0xffff, 0xffff}}
	if tm, err := DecodeTimeValue(sec); err != nil || tm.Unix() != -1 {
		t.Fatalf("seconds: %v %v", tm, err)
	}
	for _, r := range []Record{day(1), {Kind: KindValue, Codec: CodecPosixSeconds, Payload: []Word{1}}, {Kind: KindMeta, Codec: CodecGregorianDate, Payload: []Word{0, 1}}} {
		if _, err := DecodeTimeValue(r); err == nil {
			t.Fatalf("%+v: want error", r)
		}
	}
}
