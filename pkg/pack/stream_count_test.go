//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamCount — records, compressed ones and h/ shares are tallied
package pack

import "testing"

func TestStreamCount(t *testing.T) {
	s := &Stream{}
	s.count(StreamRecord{Enc: encZstd, HashOK: true})
	s.count(StreamRecord{Enc: encRaw})
	if r := s.st.Res; r.Records != 2 || r.Compressed != 1 || r.Shared != 1 {
		t.Fatalf("%+v", r)
	}
}
