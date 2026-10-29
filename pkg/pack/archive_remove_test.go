//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveRemove — stages a removal (dropping a staged put and the listing entry); absent paths are fs.ErrNotExist; remote archives fail
package pack

import (
	"context"
	"errors"
	"io/fs"
	"testing"
)

func TestArchiveRemove(t *testing.T) {
	a, _ := testArchive(t)
	a.index()
	a.WriteFile("issuers.geul", []byte("x"))
	if err := a.Remove("issuers.geul"); err != nil || !a.dels["issuers.geul"] || a.puts["issuers.geul"] != nil || hasPrefix(a.idx, "issuers.geul") {
		t.Fatal(err)
	}
	if err := a.Remove("issuers.geul"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	_, out := built(t, small())
	r, _ := OpenRemote(context.Background(), FileStore{Root: out}, Options{})
	if err := r.Remove("issuers.geul"); err == nil {
		t.Fatal("remote")
	}
}
