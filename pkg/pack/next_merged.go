//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what nextMerged — the smallest front entry over the runs (hEntryLess), advancing its run; false when all are used up
package pack

// nextMerged takes the least front entry of rs and moves that run on. With
// the few dozen runs a stream spills, a scan is enough.
func nextMerged(rs []*runReader) (hEntry, bool, error) {
	var best *runReader
	for _, rr := range rs {
		if rr.ok && (best == nil || hEntryLess(rr.cur, best.cur)) {
			best = rr
		}
	}
	if best == nil {
		return hEntry{}, false, nil
	}
	e := best.cur
	return e, true, best.advance()
}
