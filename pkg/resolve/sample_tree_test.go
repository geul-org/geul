//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what sampleTree — test fixture: A v1, A v2, B (imports A v1) and C (imports B latest, A v2) published into a temp Dir
package resolve

import (
	"testing"
)

func sampleTree(t *testing.T) Dir {
	t.Helper()
	d := Dir{Root: t.TempDir()}
	for _, b := range [][]byte{
		sampleDoc(1, "A v1"),
		sampleDoc(1, "A v2"),
		sampleDoc(2, "B", [3]uint32{1, 1, 1}),
		sampleDoc(3, "C", [3]uint32{1, 2, 0}, [3]uint32{2, 1, 2}),
	} {
		if _, err := d.Publish(b); err != nil {
			t.Fatal(err)
		}
	}
	return d
}
