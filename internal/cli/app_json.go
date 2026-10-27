//ff:func feature=cli type=helper control=sequence
//ff:what app.json — whether output is JSON, and writing a value as indented JSON
package cli

import "encoding/json"

// json writes v as indented JSON to the output and reports true when
// GEULFORMAT is json; otherwise it writes nothing and reports false.
func (a *app) json(v any) (bool, error) {
	if a.cfg.Get("GEULFORMAT") != "json" {
		return false, nil
	}
	enc := json.NewEncoder(a.out)
	enc.SetIndent("", "  ")
	return true, enc.Encode(v)
}
