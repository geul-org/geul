//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what PropFromCodec — the relation (and implied once) a record codec declares
package toulmin

// props lists every encodable relation PropFromCodec searches.
var props = []PropCode{PropAttacks, PropDo, PropUndo, PropRun, PropChecks, PropRequires, PropMeans, PropHas, PropRuns, PropEvaluates, PropTriggers, PropUntil, PropSameAs, PropImportedFrom, PropLabel, PropSupersedes}

// PropFromCodec reports the PropCode ref declares and whether the codec itself
// implies once (DO-ONCE). ok is false when ref is not a relation codec.
func PropFromCodec(ref CodecRef) (p PropCode, once bool, ok bool) {
	if ref == doCodec(true) {
		return PropDo, true, true
	}
	for _, pc := range props {
		if c, err := PropCodec(pc, false); err == nil && c == ref {
			return pc, false, true
		}
	}
	return 0, false, false
}
