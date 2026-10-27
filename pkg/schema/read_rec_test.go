//ff:func feature=schema type=rule control=sequence
//ff:what TestReadRec — prefix, untrusted and non-canonical Length, truncated and split ISSUER extension, non-canonical Issuer width
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestReadRec(t *testing.T) {
	if r, fs, ok := readRec([]core.Word{0x0882, 5}, 0, 0); !ok || fs != nil || r.Total != 2 || r.R.Payload[0] != 5 {
		t.Fatalf("plain: %+v %v %v", r, fs, ok)
	}
	if _, fs, _ := readRec([]core.Word{0x2882, 5}, 0, 0); rulesOf(fs) != "F4" || fs[0].Severity != Warn {
		t.Fatalf("001: %+v", fs)
	}
	if _, fs, _ := readRec([]core.Word{0x4882, 5}, 0, 0); rulesOf(fs) != "F4" || fs[0].Severity != Error {
		t.Fatalf("010: %+v", fs)
	}
	if _, fs, ok := readRec([]core.Word{0x0902, 5}, 0, 0); ok || rulesOf(fs) != "F2" {
		t.Fatalf("short: %+v %v", fs, ok)
	}
	if _, fs, ok := readRec([]core.Word{0x0F02, 3, 5}, 0, 0); !ok || rulesOf(fs) != "F3" {
		t.Fatalf("non-canonical: %+v %v", fs, ok)
	}
	if r, fs, ok := readRec([]core.Word{0x08FC, 1}, 0, 0); !ok || rulesOf(fs) != "F5" || len(r.R.Payload) != 1 {
		t.Fatalf("issuer truncated: %+v %v", fs, ok)
	}
	if r, fs, ok := readRec([]core.Word{0x0A7C, 1, 0xC000, 7, 8}, 0, 0); !ok || fs != nil || r.R.Issuer[0] != 1 || r.R.Format != 0xC000 || len(r.R.Payload) != 2 {
		t.Fatalf("issuer: %+v %v", r, fs)
	}
	if _, fs, ok := readRec([]core.Word{0x0A7D, 1, 0x20, 0xC000, 7}, 0, 0); !ok || fs != nil {
		t.Fatalf("2-word issuer: %+v %v", fs, ok)
	}
	if _, fs, ok := readRec([]core.Word{0x0A7D, 0, 0x20, 0xC000, 7}, 0, 0); !ok || rulesOf(fs) != "F5" {
		t.Fatalf("2-word issuer of a 1-word value: %+v %v", fs, ok)
	}
}
