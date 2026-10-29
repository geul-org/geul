//ff:func feature=pack type=writer control=sequence
//ff:what Stream.saveState — write the checkpoint (gob) to a temp file, sync it and rename it over the state file
package pack

import (
	"bytes"
	"encoding/gob"
	"path/filepath"
)

// saveState replaces out/.stream/state with the current streamState in one
// rename (writeFileSync), so a crash leaves either the old or the new
// checkpoint.
func (s *Stream) saveState() error {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s.st) // plain exported fields: never fails
	return writeFileSync(filepath.Join(s.out, streamDir, streamStateFile), b.Bytes())
}
