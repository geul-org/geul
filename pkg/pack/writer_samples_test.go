//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriterSamples — alternating train and test picks up to DefaultDictSamples; the test side stops at its cap; Get errors fail
package pack

import (
	"testing"
)

func TestWriterSamples(t *testing.T) {
	paths, get := synth(DefaultDictSamples + 10)
	w := testWriter(t, t.TempDir(), Batch{Put: paths, Get: get})
	w.plan()
	train, test, err := w.samples()
	if err != nil || len(train) != DefaultDictSamples/2 || len(test) != DefaultDictSamples/2 {
		t.Fatal(err, len(train), len(test))
	}
	big, names := map[string][]byte{}, []string{}
	for i := 0; i < 12; i++ {
		names = append(names, synthPath(i))
		big[synthPath(i)] = make([]byte, dictTrainBytes/8+i)
	}
	w = testWriter(t, t.TempDir(), Batch{Put: names, Get: memGet(big)})
	w.plan()
	if train, test, err = w.samples(); err != nil || len(test) != 2 || len(train) != 6 {
		t.Fatal(err, len(train), len(test))
	}
	w.b.Get = memGet(nil)
	if _, _, err := w.samples(); err == nil {
		t.Fatal("get")
	}
}
