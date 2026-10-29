//ff:func feature=pack type=encoder control=sequence
//ff:what TestZstdEncoder — plain and per-dictionary encoders are made once and kept; a bad dictionary fails
package pack

import "testing"

func TestZstdEncoder(t *testing.T) {
	z := &Zstd{Level: 3, Workers: 1}
	e, err := z.encoder(nil)
	if err != nil || e == nil || z.plain != e {
		t.Fatal(err)
	}
	if again, _ := z.encoder(nil); again != e {
		t.Fatal("plain not kept")
	}
	d := &Dict{ID: 9, Bytes: []byte("some raw dictionary content")}
	de, err := z.encoder(d)
	if err != nil || de == e || d.enc != de {
		t.Fatal(err)
	}
	if again, _ := z.encoder(d); again != de {
		t.Fatal("dictionary encoder not kept")
	}
	bad := []byte{0x37, 0xa4, 0x30, 0xec, 9, 0, 0, 0}
	if _, err := z.encoder(&Dict{ID: 9, Bytes: bad}); err == nil {
		t.Fatal("bad dictionary")
	}
}
