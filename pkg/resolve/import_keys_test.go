//ff:func feature=resolve type=helper control=sequence
//ff:what TestImportKeys — each IMPORT-DOC as a Key by IMPORT ID; none for a document without imports
package resolve

import (
	"github.com/geul-org/gogeul/pkg/inspect"
	"testing"
)

func TestImportKeys(t *testing.T) {
	doc, _ := inspect.Decode(sampleDoc(3, "C", [3]uint32{1, 2, 0}, [3]uint32{2, 1, 2}))
	ks := importKeys(doc)
	if len(ks) != 2 || ks[1].DocKey() != sampleKey(2, 0).DocKey() || ks[1].Version != 0 || ks[2].Version != 2 {
		t.Fatalf("got %+v", ks)
	}
	plain, _ := inspect.Decode(sampleDoc(1, "A"))
	if len(importKeys(plain)) != 0 {
		t.Fatal("want none")
	}
}
