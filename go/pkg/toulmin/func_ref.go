//ff:type feature=toulmin type=model
//ff:what FuncRef — the function slot of a `do`/`undo`/`using`: Alias REF + VerbCode
package toulmin

// FuncRef is the function slot of a `do`/`undo`/`using` (GEUL t002 TOULMIN
// NODE/LINK): a REF to the package ALIAS record (0 for a local/self-subject
// function) and a 1-word VerbCode. Verb is provisionally unbound (0x0000
// allowed) — the packet layer never interprets it. Label holds the function's
// source spelling; the stream layer writes it as a TEXT + LABEL pair, so a
// packet round-trip preserves only Alias/Verb.
type FuncRef struct {
	Alias TID
	Verb  uint16
	Label string
}
