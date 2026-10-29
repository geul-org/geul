//ff:func feature=pack type=encoder control=sequence
//ff:what hPath — `h/0x<hex>.geul`, the path of an H key
package pack

import "encoding/hex"

// hPath is the path of a 32-byte content hash key.
func hPath(h string) string {
	return "h/0x" + hex.EncodeToString([]byte(h)) + ".geul"
}
