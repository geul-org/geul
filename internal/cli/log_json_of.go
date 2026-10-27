//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what logJSONOf — a document's history as geul log JSON: hashes in hex, times in RFC 3339
package cli

import (
	"encoding/hex"
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// logJSONOf turns h into logJSON, writing each time as the text log does
// (logTime) except that an unknown time is left out.
func logJSONOf(h resolve.History) logJSON {
	out := logJSON{Versions: []logVersion{}, NextLocalID: h.Next}
	for _, e := range h.Entries {
		at := strings.TrimPrefix(logTime(e.At), "-")
		out.Versions = append(out.Versions, logVersion{Version: e.Version, At: at, Hash: "0x" + hex.EncodeToString(e.Hash[:]), State: strings.TrimPrefix(stateOf(e.VersionEntry), "\t")})
	}
	return out
}
