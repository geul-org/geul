//ff:type feature=resolve type=model
//ff:what LogEntry — one version of a document's history with the time its version list first recorded it
package resolve

import "time"

// LogEntry is a version list entry and At, the RECORDED-AT of the first list
// that carried its PUBLICATION (edge/PUBLICATION: a record's time is its first
// list's); zero when that list is no longer reachable.
type LogEntry struct {
	VersionEntry
	At time.Time
}
