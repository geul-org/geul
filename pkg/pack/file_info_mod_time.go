//ff:func feature=pack type=helper control=sequence
//ff:what fileInfo.ModTime — the zero time (pack records carry none)
package pack

import "time"

// ModTime is unknown.
func (f fileInfo) ModTime() time.Time { return time.Time{} }
