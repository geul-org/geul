//ff:func feature=cli type=command control=sequence
//ff:what TestPackVerify — one ok line per set (the top set shown as .); damage and directory archives fail
package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPackVerify(t *testing.T) {
	h := newHarness(t)
	h.doc()
	single := filepath.Join(t.TempDir(), "single")
	var out bytes.Buffer
	packBuild(h.data, single, true, &out)
	out.Reset()
	if err := packVerify(single, &out); err != nil || !strings.HasPrefix(out.String(), ". generation 1") {
		t.Fatal(err, out.String())
	}
	if err := packVerify(h.data, &out); err == nil {
		t.Fatal("directory archive")
	}
	objs, _ := filepath.Glob(filepath.Join(single, ".well-known", "geul", "o", "*.gpk"))
	os.Chmod(objs[0], 0o644)
	os.WriteFile(objs[0], []byte("damaged"), 0o644)
	if err := packVerify(single, &out); err == nil {
		t.Fatal("damaged")
	}
}
