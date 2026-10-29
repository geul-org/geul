//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what parseManifest — a `.gpm`: header (magic, version, generation, chain), section table, the six required sections
package pack

// parseManifest decodes a manifest (§4.3). Layer roots stay undecoded.
func parseManifest(b []byte) (*manifest, error) {
	c := &cursor{b: b}
	magic, ver, res, n := string(c.take(8)), c.u16(), c.u16(), c.u32()
	m := &manifest{gen: c.u64(), created: int64(c.u64()), prev: c.hash()}
	if c.bad || magic != manifestMagic || res != 0 || m.gen == 0 {
		return nil, corrupt("manifest header")
	}
	if ver != formatVersion {
		return nil, corrupt("manifest format version")
	}
	secs, err := readSections(c, n, b)
	if err != nil {
		return nil, err
	}
	for k := uint16(secObjects); k <= secHLayers; k++ {
		if _, ok := secs[k]; !ok {
			return nil, corrupt("manifest lacks a required section")
		}
	}
	ic := &cursor{b: secs[secIssuers]}
	m.issuers, err = parseIssuers(ic)
	if err == nil {
		err = ic.end("ISSUERS")
	}
	if err == nil {
		m.objects, err = parseObjects(secs[secObjects])
	}
	if err == nil {
		m.dicts, err = parseDicts(secs[secDicts])
	}
	if err == nil {
		m.named, err = parseNamed(secs[secNamed])
	}
	if err == nil {
		m.pLayers, err = parseLayers(secs[secPLayers])
	}
	if err == nil {
		m.hLayers, err = parseLayers(secs[secHLayers])
	}
	if err != nil {
		return nil, err
	}
	return m, nil
}
