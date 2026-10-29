//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirCodecSpecOf — the specs of one codec document; an unpublished codec fails
package resolve

import (
	"testing"
)

func TestDirCodecSpecOf(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.PublishCodec(codecDocOf(0x1018, 4, 2))
	s, err := d.codecSpecOf(0x1018)
	if err != nil || len(s[0x1018]) != 2 {
		t.Fatal(s, err)
	}
	if _, err := d.codecSpecOf(0x101A); err == nil {
		t.Fatal("unpublished")
	}
}
