//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirCodecSpec — specs of every published codec; stray folders skipped; a broken document fails
package resolve

import (
	"testing"
)

func TestDirCodecSpec(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.PublishCodec(codecDocOf(0x1018, 4, 2))
	d.PublishCodec(codecDocOf(0x101A, 4))
	d.write("proposal/edge/junk/versions.geul", []byte("x"))
	s, err := d.CodecSpec()
	if err != nil || len(s) != 2 || len(s[0x1018]) != 2 {
		t.Fatalf("got %+v %v", s, err)
	}
	d.write(versionPath(codecKey(0x101A), 1), codecDocOf(0x101A, 9))
	if _, err := d.CodecSpec(); err == nil {
		t.Fatal("tampered: want error")
	}
}
