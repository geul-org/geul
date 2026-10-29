//ff:func feature=pack type=writer control=sequence
//ff:what openPart — reopen a stream's open object at its checkpointed size (nil when none was open)
package pack

import (
	"fmt"
	"os"
	"path/filepath"
)

// openPart opens the part file of an object open at the checkpoint and cuts
// it back to open.Off, dropping bytes written after the checkpoint. A part
// shorter than that was not synced as the checkpoint says: an error.
func openPart(out, part string, open streamOpen) (*os.File, error) {
	if open.Num < 0 {
		return nil, nil
	}
	f, err := os.OpenFile(filepath.Join(out, streamDir, part), os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err == nil && uint64(fi.Size()) < open.Off {
		err = fmt.Errorf("%w: %s holds %d bytes, the checkpoint %d", ErrCorrupt, part, fi.Size(), open.Off)
	}
	if err == nil {
		err = f.Truncate(int64(open.Off))
	}
	if err != nil {
		f.Close()
		return nil, err
	}
	return f, nil
}
