//ff:func feature=pack type=writer control=sequence
//ff:what Stream.Close — finish the last data pack, the P and H layers and the index objects, write the manifest and latest.txt, remove the working folder
package pack

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Close ends the stream (§6 steps 4–6): the open data pack gets its
// inventory, the P layer (pLayer) and the H layer (hLayer) are finished, the
// open index object is closed, and the manifest of generation 1 (no previous
// manifest) is written, then latest.txt, last. The working folder is removed.
// A stream with no record is an error.
func (s *Stream) Close() (BuildResult, error) {
	if s.st.Count == 0 && len(s.st.Named) == 0 {
		return BuildResult{}, errors.New("geul/pack: a stream with no record makes no set")
	}
	named, err := sortedNamed(s.st.Named)
	if err == nil && s.st.Data.Num >= 0 {
		err = s.finishData()
	}
	var pl, hl []layer
	if err == nil {
		pl, err = s.pLayer()
	}
	if err == nil {
		hl, err = s.hLayer()
	}
	if err == nil && s.st.Index.Num >= 0 {
		err = s.closeObject(s.index, s.st.Index, kindIndex, nil)
	}
	if err != nil {
		return BuildResult{}, err
	}
	mb := encodeManifest(s.manifest(named, pl, hl))
	sum := sha256.Sum256(mb)
	name := hex.EncodeToString(sum[:])
	if err := writeFileSync(filepath.Join(s.out, "m", name+".gpm"), mb); err != nil {
		return BuildResult{}, err
	}
	if err := writeFileSync(filepath.Join(s.out, latestFile), []byte(fmt.Sprintf("%d %s\n", s.st.Gen, name))); err != nil {
		return BuildResult{}, err
	}
	res := s.st.Res
	res.Generation, res.Manifest, res.Changed, res.Bytes = s.st.Gen, name, true, res.Bytes+int64(len(mb))
	return res, os.RemoveAll(filepath.Join(s.out, streamDir))
}
