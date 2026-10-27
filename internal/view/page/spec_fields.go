//ff:func feature=page type=builder control=iteration dimension=1
//ff:what specFields — a record's declared payload fields as shown fields; raw fields only when it has no spec
package page

import "github.com/geul-org/gogeul/pkg/inspect"

// specFields shows each declared payload field of e (edge/DECLARE), the Edge
// Id aside, by name and partText. A record without a codec spec shows
// rawFields.
func specFields(e inspect.Entry) []Field {
	if e.Spec == nil {
		return rawFields(e.Record)
	}
	var fs []Field
	for _, p := range inspect.Layout(e) {
		if p.Head || p.Name == "ID" {
			continue
		}
		fs = append(fs, Field{Key: p.Name, Text: partText(p)})
	}
	return fs
}
