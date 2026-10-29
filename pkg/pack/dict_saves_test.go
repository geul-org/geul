//ff:func feature=pack type=writer control=sequence
//ff:what TestDictSaves — a dictionary trained on similar documents saves on others; no test records save nothing; a dictionary that cannot encode does not save
package pack

import "testing"

func TestDictSaves(t *testing.T) {
	z := &Zstd{DictSize: 8192}
	o := streamTestOptions()
	d, err := z.Train(o.Samples[:20], FirstDictID)
	if err != nil {
		t.Fatal(err)
	}
	if !dictSaves(z, d, o.Samples[20:]) {
		t.Fatal("similar documents: want saving")
	}
	if dictSaves(z, d, nil) {
		t.Fatal("no test records: want false")
	}
	if dictSaves(z, &Dict{ID: 7, Bytes: d.Bytes}, o.Samples[20:]) {
		t.Fatal("dictionary carrying another ID: want false")
	}
}
