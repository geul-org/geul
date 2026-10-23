//ff:func feature=resolve type=resolver control=sequence
//ff:what TestFetchVersion — a listed version verifies with its Source name; a missing file and a wrong hash fail
package resolve

import (
	"context"
	"errors"
	"testing"
)

func TestFetchVersion(t *testing.T) {
	ctx := context.Background()
	tree := sampleTree(t)
	list, _ := tree.readList(sampleKey(1, 0).DocKey())
	e := Entry{Name: "tree", Source: tree}
	if n, err := fetchVersion(ctx, e, sampleKey(1, 0), list[0]); err != nil || n.Version != 1 || n.Source != "tree" {
		t.Fatalf("ok: %+v %v", n, err)
	}
	if _, err := fetchVersion(ctx, Entry{Source: Memory{}}, sampleKey(1, 0), list[0]); !errors.Is(err, ErrNotFound) {
		t.Fatalf("missing: %v", err)
	}
	if _, err := fetchVersion(ctx, e, sampleKey(1, 0), VersionEntry{Version: 1}); !errors.Is(err, ErrVerify) {
		t.Fatalf("hash: %v", err)
	}
}
