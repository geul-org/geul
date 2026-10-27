//ff:func feature=schema type=decoder control=iteration dimension=1
//ff:what statementsOf — the TRIPLEs and QUALIFIERs of a BODY in document order
package schema

// statementsOf keeps every BODY record splitStatement cuts.
func statementsOf(body []rec) []statement {
	var out []statement
	for _, r := range body {
		if st, ok := splitStatement(r); ok {
			out = append(out, st)
		}
	}
	return out
}
