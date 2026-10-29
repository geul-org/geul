//ff:func feature=pack type=writer control=sequence
//ff:what writer.checkLatest — ErrConflict when latest.txt no longer names the generation this write started from
package pack

import (
	"fmt"
	"os"
	"path/filepath"
)

// checkLatest guards against a concurrent writer before latest.txt moves.
func (w *writer) checkLatest() error {
	want := ""
	if w.prev != nil {
		want = fmt.Sprintf("%d %s\n", w.prev.m.gen, w.prev.name)
	}
	b, _ := os.ReadFile(filepath.Join(w.out, latestFile))
	if string(b) != want {
		return ErrConflict
	}
	return nil
}
