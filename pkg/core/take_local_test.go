//ff:func feature=core type=decoder control=sequence
//ff:what TestTakeLocal — a LOCAL record's ID and size; other records and junk fail
package core

import (
	"testing"
)

func TestTakeLocal(t *testing.T) {
	if id, n, err := takeLocal(append(localWords(7), 9)); err != nil || id != 7 || n != 2 {
		t.Fatalf("got %d %d %v", id, n, err)
	}
	if _, _, err := takeLocal(nil); err == nil {
		t.Fatal("empty: want error")
	}
	ws, _ := EncodeRecord(EncodeUintValue(1))
	if _, _, err := takeLocal(ws); err == nil {
		t.Fatal("value: want error")
	}
}
