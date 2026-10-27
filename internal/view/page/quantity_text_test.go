//ff:func feature=page type=helper control=sequence
//ff:what TestQuantityText — amount alone, with unit and bounds, and with one bound missing
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestQuantityText(t *testing.T) {
	if got := quantityText(inspect.Value{Amount: "1.20"}); got != "1.20" {
		t.Fatalf("alone: %q", got)
	}
	if got := quantityText(inspect.Value{Amount: "652", Unit: "Q828224", Lower: "651", Upper: "653"}); got != "652 Q828224 [651, 653]" {
		t.Fatalf("full: %q", got)
	}
	if got := quantityText(inspect.Value{Amount: "5", Upper: "6"}); got != "5 [?, 6]" {
		t.Fatalf("one bound: %q", got)
	}
}
