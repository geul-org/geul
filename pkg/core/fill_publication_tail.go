//ff:func feature=core type=decoder control=selection
//ff:what fillPublicationTail — copy the content hash or the reason from a version list record's tail
package core

// fillPublicationTail stores t, already of the codec's tail length, into p: the
// 16-word content hash for PUBLICATION, the reason word for TOMBSTONE.
func fillPublicationTail(p *Publication, t []Word) {
	switch p.Codec {
	case EdgePublication:
		copy(p.Hash[:], PutWords(t))
	case EdgeTombstone:
		p.Reason = uint16(t[0])
	}
}
