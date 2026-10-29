//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterDictPays — a useful dictionary pays on many records, not on few; no samples and encoder errors do not pay
package pack

import "testing"

func TestWriterDictPays(t *testing.T) {
	d, samples := testDict(t, FirstDictID, 200)
	paths, get := synth(200)
	w := testWriter(t, t.TempDir(), Batch{Put: paths, Get: get})
	w.plan()
	z := &Zstd{Workers: 1}
	if !w.dictPays(z, d, samples[:50]) {
		t.Fatal("200 records")
	}
	w2 := testWriter(t, t.TempDir(), Batch{Put: paths[:2], Get: get})
	w2.plan()
	if w2.dictPays(z, d, samples[:50]) {
		t.Fatal("2 records")
	}
	if w.dictPays(z, d, nil) {
		t.Fatal("no samples")
	}
	if w.dictPays(z, &Dict{ID: 1, Bytes: d.Bytes}, samples[:5]) {
		t.Fatal("dictionary encoder")
	}
}
