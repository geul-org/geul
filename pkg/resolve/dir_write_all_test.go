//ff:func feature=resolve type=rule control=iteration dimension=1
//ff:what TestDirWriteAll — every path written, the list last; a failing write stops
package resolve

import (
	"os"
	"testing"
)

func TestDirWriteAll(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	files := map[string][]byte{"a/versions.geul": []byte("l"), "a/b.geul": []byte("b"), "h/x.geul": []byte("x")}
	if err := d.writeAll(files); err != nil {
		t.Fatal(err)
	}
	for p, want := range files {
		if b, _ := os.ReadFile(d.file(p)); string(b) != string(want) {
			t.Fatalf("%s: %q", p, b)
		}
	}
	blockWrite(t, d, "z/b.geul")
	if err := d.writeAll(map[string][]byte{"z/b.geul": nil, "z/versions.geul": nil}); err == nil {
		t.Fatal("blocked: want error")
	}
	if _, err := os.Stat(d.file("z/versions.geul")); err == nil {
		t.Fatal("list written after a failed write")
	}
}
