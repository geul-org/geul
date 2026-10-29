//ff:func feature=pack type=rule control=sequence
//ff:what TestArchiveVerify — every set passes; a damaged set fails naming it; sets without a generation are skipped
package pack

import (
	"context"
	"strings"
	"testing"
)

func TestArchiveVerify(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	a.set("0x0001")
	if err := a.Verify(ctx); err != nil {
		t.Fatal(err)
	}
	flip(t, objPath(t, a.sets[0].r, a.sets[0].dir, kindData), 200)
	if err := a.Verify(ctx); err == nil || !strings.Contains(err.Error(), `"0x0000"`) {
		t.Fatal(err)
	}
}
