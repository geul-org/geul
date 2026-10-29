//ff:func feature=pack type=writer control=sequence
//ff:what NewStream — start the first generation of a pack set in out for a Stream; out must hold no set and no stream
package pack

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// NewStream starts writing the first generation of the pack set in folder
// out (`latest.txt`, `m/`, `o/`), with its working files in out/.stream. A
// folder that already holds a set (latest.txt) is refused: a stream writes a
// whole set, it does not add a generation; one that holds a stream is
// refused too (ResumeStream continues it). Codec nil is Zstd at its
// defaults. When the codec trains dictionaries and o.Samples are enough, a
// dictionary is trained first and stored as the first record (streamDictOf).
func NewStream(out string, o StreamOptions) (*Stream, error) {
	o = o.withDefaults()
	if _, err := os.Stat(filepath.Join(out, latestFile)); err == nil {
		return nil, fmt.Errorf("geul/pack: %s already holds a pack set; a stream writes a new one", out)
	}
	if _, err := os.Stat(filepath.Join(out, streamDir, streamStateFile)); err == nil {
		return nil, errors.New("geul/pack: " + out + " holds an unfinished stream; resume it")
	}
	if err := makeStreamDirs(out); err != nil {
		return nil, err
	}
	s := &Stream{out: out, o: o, inum: issuerIndex(o.Issuers)}
	s.st = streamState{Gen: 1, Created: o.Created, Issuers: o.Issuers, Data: streamOpen{Num: -1}, Index: streamOpen{Num: -1}}
	if s.st.Created == 0 {
		s.st.Created = o.Build.Now().Unix()
	}
	if d := streamDictOf(o); d != nil {
		return s, s.storeDict(d)
	}
	return s, nil
}
