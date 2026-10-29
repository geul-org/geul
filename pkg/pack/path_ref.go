//ff:type feature=pack type=model
//ff:what pathRef — an archive path mapped to its key space: P (Issuer, document, slot), H (content hash) or named
package pack

// pathRef is a path's place in the key spaces (§3.4). For spaceP issuer is the
// canonical Issuer hex; for spaceH hash is the 32 raw bytes.
type pathRef struct {
	space  int
	issuer string
	doc    uint64
	slot   uint32
	hash   string
}
