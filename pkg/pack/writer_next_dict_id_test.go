//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterNextDictID — FirstDictID for a new set, one past the largest ID otherwise
package pack

import "testing"

func TestWriterNextDictID(t *testing.T) {
	if (&writer{}).nextDictID() != FirstDictID {
		t.Fatal("new")
	}
	prev := newReader(memStore{}, Options{}, &manifest{dicts: []dictEntry{{id: FirstDictID + 4}, {id: FirstDictID + 1}}}, "")
	if id := (&writer{prev: prev}).nextDictID(); id != FirstDictID+5 {
		t.Fatal(id)
	}
}
