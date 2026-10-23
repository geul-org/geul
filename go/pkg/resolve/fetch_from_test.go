//ff:func feature=resolve type=resolver control=sequence
//ff:what TestFetchFrom — latest and pinned versions with Source name; missing list, bad list, missing version, missing file and bad bytes fail
package resolve

import (
	"context"
	"errors"
	"testing"
)

func TestFetchFrom(t *testing.T) {
	ctx := context.Background()
	e := Entry{Name: "tree", Source: sampleTree(t)}
	if n, err := fetchFrom(ctx, e, sampleKey(1, 0)); err != nil || n.Version != 2 || n.Source != "tree" {
		t.Fatalf("latest: %+v %v", n, err)
	}
	if n, err := fetchFrom(ctx, e, sampleKey(1, 1)); err != nil || n.Version != 1 {
		t.Fatalf("pinned: %+v %v", n, err)
	}
	if _, err := fetchFrom(ctx, e, sampleKey(9, 0)); !errors.Is(err, ErrNotFound) {
		t.Fatalf("no list: %v", err)
	}
	if _, err := fetchFrom(ctx, e, sampleKey(1, 7)); !errors.Is(err, ErrNotFound) {
		t.Fatalf("no version: %v", err)
	}
	k := sampleKey(1, 0).DocKey()
	m := Memory{listPath(k): []byte("junk")}
	if _, err := fetchFrom(ctx, Entry{Source: m}, sampleKey(1, 0)); !errors.Is(err, ErrVerify) {
		t.Fatalf("bad list: %v", err)
	}
	m[listPath(k)] = []byte(formatVersionList([]VersionEntry{{Version: 1}}))
	if _, err := fetchFrom(ctx, Entry{Source: m}, sampleKey(1, 0)); !errors.Is(err, ErrNotFound) {
		t.Fatalf("no file: %v", err)
	}
	m[versionPath(k, 1)] = sampleDoc(1, "A v1")
	if _, err := fetchFrom(ctx, Entry{Source: m}, sampleKey(1, 0)); !errors.Is(err, ErrVerify) {
		t.Fatalf("bad hash: %v", err)
	}
}
