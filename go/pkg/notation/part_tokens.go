//ff:func feature=notation type=decoder control=selection
//ff:what partTokens — a part as tokens: a generic payload that is a run of records with short forms record by record, any other part as one token
package notation

import "github.com/geul-org/geul/go/pkg/inspect"

// partTokens writes p (formatPart). A generic "payload" part (no declared
// layout, read back token by token by parseGeneric) whose words are wholly
// records that all have a short form is written one token per record
// (recordTokens), which reads back to the same words.
func partTokens(p inspect.Part) []string {
	switch {
	case p.Name == "payload" && !p.Record:
		if ts, ok := recordTokens(p.Words); ok {
			return ts
		}
	}
	return []string{formatPart(p)}
}
