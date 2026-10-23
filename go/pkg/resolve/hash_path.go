//ff:func feature=resolve type=helper control=sequence
//ff:what hashPath — `h/<내용 해시>.geul`, the address by content hash alone
package resolve

import "encoding/hex"

// hashPath is the immutable address of a document by its content hash alone.
func hashPath(h [32]byte) string {
	return "h/0x" + hex.EncodeToString(h[:]) + ".geul"
}
