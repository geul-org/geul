//ff:func feature=pack type=writer control=sequence
//ff:what Stream.startObject — reserve the next object number and open its part file, items starting after the 64-byte header
package pack

import (
	"os"
	"path/filepath"
)

// startObject reserves the next OBJECTS position for an object of kind (the
// position is its number; data packs keep their start order, §4.3) and
// creates its part file in the working folder. Items are written from
// offset 64; closeObject fills the header in front of them.
func (s *Stream) startObject(kind byte, part string) (*os.File, streamOpen, error) {
	f, err := os.OpenFile(filepath.Join(s.out, streamDir, part), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return nil, streamOpen{}, err
	}
	s.st.Objects = append(s.st.Objects, streamObject{Kind: kind})
	return f, streamOpen{Num: len(s.st.Objects) - 1, Off: headerSize}, nil
}
