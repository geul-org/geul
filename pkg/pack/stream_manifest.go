//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.manifest — the manifest of the stream's generation from its objects, Issuers, dictionary, named records and layers
package pack

// manifest assembles §4.3: generation 1 with no previous manifest, the
// reserved objects in order, the Issuers, the dictionary (if any), NAMED and
// the two layers.
func (s *Stream) manifest(named []namedEntry, pl, hl []layer) *manifest {
	m := &manifest{gen: s.st.Gen, created: s.st.Created, issuers: s.st.Issuers, named: named, pLayers: pl, hLayers: hl}
	for _, o := range s.st.Objects {
		m.objects = append(m.objects, object{kind: o.Kind, size: o.Size, hash: o.Hash})
	}
	if s.st.Dict.ID != 0 {
		m.dicts = []dictEntry{{id: s.st.Dict.ID, loc: s.st.Dict.Loc, hash: s.st.Dict.Hash}}
	}
	return m
}
