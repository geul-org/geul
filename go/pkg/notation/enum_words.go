//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what enumWords — a raw field's words from an ENUM name (case ignored) or raw words
package notation

import (
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
)

// enumWords returns the code whose ENUM name is tok (enums maps a code in
// canonical hex to its name), else tok read as raw words (parseHex).
func enumWords(tok string, enums map[string]string) ([]core.Word, error) {
	for code, name := range enums {
		if strings.EqualFold(name, tok) {
			return parseHex(code)
		}
	}
	return parseHex(tok)
}
