//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirCodecSpec — specs of every published codec; stray folders, folders without a version list and other GEUL documents skipped; a broken document fails
package resolve

import (
	"testing"
)

func TestDirCodecSpec(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.PublishCodec(codecDocOf(0x1018, 4, 2))
	d.PublishCodec(codecDocOf(0x101A, 4))
	d.write("0x0000/junk/versions.geul", []byte("x"))
	d.write("0x0000/0xffff000000000001/versions.geul", []byte("x"))
	d.write(versionPath(codecKey(0x101C), 1), []byte("no list"))
	s, err := d.CodecSpec()
	if err != nil || len(s) != 2 || len(s[0x1018]) != 2 {
		t.Fatalf("got %+v %v", s, err)
	}
	d.write(versionPath(codecKey(0x101A), 1), codecDocOf(0x101A, 9))
	if _, err := d.CodecSpec(); err == nil {
		t.Fatal("tampered: want error")
	}
}
