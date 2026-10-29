//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.writeBlocks — store blocks as checksummed zstd frames in index objects, opening a new one when full
package pack

// writeBlocks writes bs and returns their locations.
func (w *writer) writeBlocks(bs []block) ([]Loc, error) {
	locs := make([]Loc, len(bs))
	for i, b := range bs {
		frame, err := w.o.Codec.Compress(encodeBlock(b), nil)
		if err == nil && w.cur != nil && !w.cur.fits(len(frame), w.o.PackSize) {
			err = w.closeObject(nil)
		}
		if err == nil && w.cur == nil {
			err = w.startObject(kindIndex)
		}
		if err != nil {
			return nil, err
		}
		off := w.cur.add(frame)
		locs[i] = Loc{Obj: w.curNum, Off: off, Len: uint64(len(frame)), Enc: encZstd}
	}
	return locs, nil
}
