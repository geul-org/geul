//ff:func feature=pack type=writer control=sequence
//ff:what Stream.storeDict — store d as the first record of the first data pack and make it the dictionary records use
package pack

import "crypto/sha256"

// storeDict places d's bytes as an inventory kind 4 record (§4.5) and records
// it for DICTS and for the checkpoint.
func (s *Stream) storeDict(d *Dict) error {
	loc, err := s.place(d.Bytes, encRaw)
	if err != nil {
		return err
	}
	s.st.Inv = append(s.st.Inv, streamInv{Kind: invDict, Ref: d.ID, N: uint32(len(d.Bytes)), Off: loc.Off})
	s.st.Dict = streamDict{ID: d.ID, Loc: loc, Hash: sha256.Sum256(d.Bytes)}
	s.st.DictRaw = d.Bytes
	s.st.Res.Dict = d.ID
	s.dict = d
	return nil
}
