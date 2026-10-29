//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamFinishData — the pack closes with its inventory and the next starts empty; an inventory that does not compress fails
package pack

import "testing"

func TestStreamFinishData(t *testing.T) {
	s := streamTestNew(t, nil)
	s.place([]byte("x"), encRaw)
	s.st.Inv = []streamInv{{Kind: invNamed, N: 1, Off: headerSize}}
	s.st.Paths = []string{"a.geul"}
	if err := s.finishData(); err != nil || s.st.Data.Num != -1 || s.st.Inv != nil || s.st.Paths != nil || s.st.Objects[0].Size == 0 {
		t.Fatalf("%+v %v", s.st.Data, err)
	}
	s2 := streamTestNew(t, failCodec{})
	s2.place([]byte("x"), encRaw)
	if err := s2.finishData(); err == nil {
		t.Fatal("failing codec: want error")
	}
}
