//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestPlaceDocRecord — HEAD/BODY/END routing and every placement error
package core

import "testing"

func TestPlaceDocRecord(t *testing.T) {
	st := &docState{}
	meta := Record{Kind: KindMeta, Codec: MetaProfile}
	edge := Record{Kind: KindEdge, Codec: 0x60}
	for _, r := range []Record{meta, {Kind: KindMeta, Codec: MetaBody}, edge, {Kind: KindMeta, Codec: MetaEnd}} {
		if err := placeDocRecord(st, r); err != nil {
			t.Fatal(err)
		}
	}
	if len(st.doc.Head) != 1 || len(st.doc.Body) != 1 || !st.ended {
		t.Fatalf("routing: %+v", st)
	}
	if err := placeDocRecord(st, edge); err == nil {
		t.Fatal("record after END must fail")
	}
	implicit := &docState{}
	if err := placeDocRecord(implicit, edge); err != nil || !implicit.inBody {
		t.Fatal("non-Meta first record starts an implicit BODY")
	}
	if err := placeDocRecord(implicit, Record{Kind: KindMeta, Codec: MetaBody}); err == nil {
		t.Fatal("BODY after body records must fail")
	}
	if err := placeDocRecord(implicit, meta); err == nil {
		t.Fatal("Meta in BODY must fail")
	}
	if err := placeDocRecord(&docState{}, Record{Kind: KindMeta, Codec: MetaDoc1Word}); err == nil {
		t.Fatal("nested DOC must fail")
	}
	head := &docState{}
	_ = placeDocRecord(head, meta)
	if err := placeDocRecord(head, edge); err == nil {
		t.Fatal("non-Meta in HEAD must fail")
	}
}
