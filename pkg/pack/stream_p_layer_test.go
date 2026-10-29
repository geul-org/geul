//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamPLayer — no record no layer; one leaf is the root; more get a root over every written leaf; a failing leaf write fails
package pack

import "testing"

func TestStreamPLayer(t *testing.T) {
	if ls, err := streamTestNew(t, nil).pLayer(); ls != nil || err != nil {
		t.Fatal("no record")
	}
	one := streamTestNew(t, nil)
	streamTestPut(t, one, 1, 2)
	if ls, err := one.pLayer(); err != nil || ls[0].count != 4 || one.st.Index.Num != -1 {
		t.Fatal("single leaf is the root")
	}
	many := streamTestNew(t, nil)
	streamTestPut(t, many, 1, 6)
	ls, err := many.pLayer()
	if err != nil || ls[0].count != 12 || len(many.st.PHeads) != 3 {
		t.Fatalf("%v %d", err, len(many.st.PHeads))
	}
	bad := streamTestNew(t, nil)
	streamTestPut(t, bad, 1, 6)
	bad.o.Build.Codec = failCodec{}
	if _, err := bad.pLayer(); err == nil {
		t.Fatal("failing leaf write: want error")
	}
}
