//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamClose — a closed stream is generation 1 with latest.txt last and no working folder; no record, a named path twice, a failing data pack, P or H layer, index object, manifest or latest.txt write is an error
package pack

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStreamClose(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 6)
	res, err := s.Close()
	if err != nil || res.Generation != 1 || !res.Changed || res.Records != 12 || res.Manifest == "" {
		t.Fatalf("%+v %v", res, err)
	}
	if b, _ := os.ReadFile(filepath.Join(s.out, latestFile)); !strings.HasPrefix(string(b), "1 "+res.Manifest) {
		t.Fatal(string(b))
	}
	if _, err := os.Stat(filepath.Join(s.out, streamDir)); err == nil {
		t.Fatal("working folder left")
	}
	if _, err := streamTestNew(t, nil).Close(); err == nil {
		t.Fatal("no record: want error")
	}
	twice := streamTestNew(t, nil)
	twice.st.Named = []streamKV{{Key: "a"}, {Key: "a"}}
	if _, err := twice.Close(); err == nil {
		t.Fatal("named twice: want error")
	}
	cases := map[string]func(*Stream){
		"data pack":    func(s *Stream) { s.o.Build.Codec = failCodec{} },
		"P layer":      func(s *Stream) { s.finishData(); s.o.Build.Codec = failCodec{} },
		"H layer":      func(s *Stream) { s.st.Runs = []string{"run-missing"} },
		"index object": func(s *Stream) { s.flushLeaf(); s.index.Close() },
		"manifest":     func(s *Stream) { os.RemoveAll(filepath.Join(s.out, "m")) },
		"latest.txt":   func(s *Stream) { os.MkdirAll(filepath.Join(s.out, latestFile, "x"), 0o755) },
	}
	for name, breakIt := range cases {
		b := streamTestNew(t, nil)
		streamTestPut(t, b, 1, 6)
		breakIt(b)
		if _, err := b.Close(); err == nil {
			t.Errorf("%s: want error", name)
		}
	}
}
