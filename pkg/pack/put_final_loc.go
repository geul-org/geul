//ff:func feature=pack type=helper control=sequence
//ff:what put.finalLoc — the record that serves a put: its own, or the shared record's
package pack

// finalLoc is where p's bytes ended up.
func (p *put) finalLoc() Loc {
	if p.via != nil {
		return p.via.loc
	}
	return p.loc
}
