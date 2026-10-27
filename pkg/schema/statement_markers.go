//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what statementMarkers — E7: slot markers in TRIPLEs and QUALIFIERs only as Object/Value UNKNOWN-VALUE or NO-VALUE
package schema

// statementMarkers applies E7 to every TRIPLE and QUALIFIER of body
// (markerFacts): the grammar fixes their fields, and only the Object
// (TRIPLE) and Value (QUALIFIER) take the slot markers UNKNOWN-VALUE and
// NO-VALUE (edge/TRIPLE, QUALIFIER). Other codecs' fields are checked
// against their spec documents, not here.
func statementMarkers(body []rec) []Fact {
	var fs []Fact
	for _, s := range statementsOf(body) {
		fs = append(fs, markerFacts(s)...)
	}
	return fs
}
