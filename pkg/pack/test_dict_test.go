//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what testDict — a dictionary trained on synth records, and those records
package pack

import "testing"

// testDict trains a small dictionary with id on synth(n).
func testDict(t *testing.T, id uint32, n int) (*Dict, [][]byte) {
	t.Helper()
	paths, get := synth(n)
	var samples [][]byte
	for _, p := range paths {
		b, _ := get(p)
		samples = append(samples, b)
	}
	d, err := (&Zstd{DictSize: 4096}).Train(samples, id)
	if err != nil {
		t.Fatal(err)
	}
	return d, samples
}
