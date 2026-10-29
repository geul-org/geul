//ff:func feature=pack type=writer control=sequence
//ff:what TestRunReaderAdvance — whole entries load, the end clears ok, a cut entry is corrupt
package pack

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestRunReaderAdvance(t *testing.T) {
	e := hEntry{loc: Loc{Obj: 1, Off: 2}}
	rr := &runReader{f: os.Stdin, r: bufio.NewReader(bytes.NewReader(appendHEntry(nil, e)))}
	if err := rr.advance(); err != nil || !rr.ok || rr.cur != e {
		t.Fatal(err)
	}
	if err := rr.advance(); err != nil || rr.ok {
		t.Fatal("end")
	}
	cut := &runReader{f: os.Stdin, r: bufio.NewReader(bytes.NewReader(make([]byte, 10)))}
	if err := cut.advance(); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
