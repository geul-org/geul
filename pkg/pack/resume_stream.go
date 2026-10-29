//ff:func feature=pack type=writer control=sequence
//ff:what ResumeStream — continue the stream in out from its last checkpoint: state, open part files cut back to their checkpointed size, stray runs removed
package pack

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"
)

// ResumeStream reopens the unfinished stream in out at its last Checkpoint.
// The open data pack and index object are cut back to the bytes the
// checkpoint counted, runs spilled after it are removed, and Count and Last
// tell the caller where to go on. o must carry the same codec settings as the
// stream was started with (the dictionary and Issuers come from the
// checkpoint), or the output differs from an uninterrupted run.
func ResumeStream(out string, o StreamOptions) (*Stream, error) {
	b, err := os.ReadFile(filepath.Join(out, streamDir, streamStateFile))
	if err != nil {
		return nil, err
	}
	s := &Stream{out: out, o: o.withDefaults()}
	if err := gob.NewDecoder(bytes.NewReader(b)).Decode(&s.st); err != nil {
		return nil, err
	}
	s.inum = issuerIndex(s.st.Issuers)
	if s.st.Dict.ID != 0 {
		s.dict = &Dict{ID: s.st.Dict.ID, Bytes: s.st.DictRaw}
	}
	if s.data, err = openPart(out, streamDataPart, s.st.Data); err != nil {
		return nil, err
	}
	if s.index, err = openPart(out, streamIndexPart, s.st.Index); err != nil {
		return nil, err
	}
	return s, removeStrayRuns(out, s.st.Runs)
}
