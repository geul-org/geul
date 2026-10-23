//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirPublishCodec — a spec document lands at its codec address; odd bytes, bad hash input and documents without one DECLARE fail
package resolve

import (
	"github.com/geul-org/geul/go/pkg/core"
	"os"
	"testing"
)

func TestDirPublishCodec(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	if e, err := d.PublishCodec(codecDocOf(0x1018, 4, 2)); err != nil || e.Version != 1 {
		t.Fatalf("publish: %+v %v", e, err)
	}
	if _, err := os.Stat(d.file("proposal/edge/0011000/0x00000001.geul")); err != nil {
		t.Fatal(err)
	}
	if _, err := d.PublishCodec([]byte{1}); err == nil {
		t.Fatal("odd bytes: want error")
	}
	if _, err := d.PublishCodec(core.PutWords(core.WrapDocument([]core.Word{0x180E}))); err == nil {
		t.Fatal("no END: want error")
	}
	if _, err := d.PublishCodec(core.PutWords(core.WrapDocument([]core.Word{0x180E, 0x180F}))); err == nil {
		t.Fatal("no DECLARE: want error")
	}
}
