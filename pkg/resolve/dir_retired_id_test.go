//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirRetiredID — a dropped ID is refused by Publish; IDs from NEXT-LOCAL-ID on, a diff and a first version pass; a broken stored version fails
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/delta"
)

func TestDirRetiredID(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	k := sampleKeyOf(7)
	if _, err := d.Publish(edgeDoc(7, "a", "b")); err != nil {
		t.Fatal(err)
	}
	if _, err := d.Publish(edgeDoc(7, "a")); err != nil {
		t.Fatal(err)
	}
	if _, err := d.Publish(edgeDoc(7, "a", "B")); err == nil {
		t.Fatal("#2 dropped in version 2: want error")
	}
	lg, _ := d.readList(k, k.Doc)
	if err := d.retiredID(k, lg, edgeDoc(7, "a", "b", "c")); err == nil {
		t.Fatal("#2 again: want error")
	}
	diff, _ := delta.Diff(edgeDoc(7, "a"), edgeDoc(7, "a", "b"), 2)
	if err := d.retiredID(k, lg, diff); err != nil {
		t.Fatalf("diff: %v", err)
	}
	if err := d.retiredID(k, versionLog{}, edgeDoc(7, "a", "b")); err != nil {
		t.Fatalf("first version: %v", err)
	}
	d.write(versionPath(k, 2), []byte{1})
	if err := d.retiredID(k, lg, edgeDoc(7, "a")); err == nil {
		t.Fatal("broken stored version: want error")
	}
}
