//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirBasedOn — finds the version stored as a diff on v; none, a deleted one and a broken list give false
package resolve

import (
	"os"
	"testing"

	"github.com/geul-org/gogeul/pkg/delta"
)

func TestDirBasedOn(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	k := sampleKeyOf(7)
	v1, v2 := edgeDoc(7, "a"), edgeDoc(7, "a", "b")
	d.Publish(v1)
	d.Publish(v2)
	if _, ok := d.basedOn(k, 1); ok {
		t.Fatal("snapshots: want false")
	}
	diff, _ := delta.Diff(v1, v2, 1)
	os.WriteFile(d.file(versionPath(k, 2)), diff, 0o644)
	if v, ok := d.basedOn(k, 1); !ok || v != 2 {
		t.Fatalf("%d %v", v, ok)
	}
	d.write(listPath(k), []byte{1})
	if _, ok := d.basedOn(k, 1); ok {
		t.Fatal("broken list: want false")
	}
}
