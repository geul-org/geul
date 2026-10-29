//ff:type feature=pack type=model
//ff:what writer — the state of one generation being written: previous generation, plan, open object, OBJECTS so far
package pack

import "context"

// writer builds one generation in out. base is the previous live state
// (path → Loc, empty when Fresh), remap renumbers carried locations when
// OBJECTS is rebuilt, cur is the one object open at a time. o.Codec is never
// nil once open ran (writeTo resolves it).
type writer struct {
	ctx     context.Context
	out     string
	o       BuildOptions
	fs      fileSys
	b       Batch
	prev    *Reader
	base    map[string]Loc
	issuers []string
	inum    map[string]uint32
	objs    []object
	remap   map[uint32]uint32
	all     bool // every layer is rewritten this generation
	sums    map[string][32]byte
	puts    []*put
	dels    []string
	cur     *objWriter
	curNum  uint32
	inv     inventory
	dicts   []dictEntry // DICTS of the new generation: carried, then the one trained now
	dict    *Dict       // the dictionary new records use (nil: none)
	res     BuildResult
}
