//ff:func feature=resolve type=decoder control=selection
//ff:what parseEntry — one GEULPROXY item: https://, file://, registry or off
package resolve

import (
	"fmt"
	"strings"
)

// parseEntry reads one GEULPROXY item. Plain http:// needs GEULINSECURE, which
// is not supported yet, so it is rejected.
func parseEntry(s string) (Entry, error) {
	switch {
	case s == "off":
		return Entry{Name: s, Off: true}, nil
	case s == "registry":
		return Entry{Name: s, Source: Registry{}, Public: true}, nil
	case strings.HasPrefix(s, "file://") && len(s) > len("file://"):
		return Entry{Name: s, Source: Dir{Root: strings.TrimPrefix(s, "file://")}}, nil
	case strings.HasPrefix(s, "https://") && len(s) > len("https://"):
		base := strings.TrimSuffix(s, "/")
		return Entry{Name: s, Source: HTTP{Base: base}, Public: base == geulOrg}, nil
	default:
		return Entry{}, fmt.Errorf("geul/resolve: bad GEULPROXY item %q", s)
	}
}
