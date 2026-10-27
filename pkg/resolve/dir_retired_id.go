//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.retiredID — whether a document about to be published gives a new statement an ID below NEXT-LOCAL-ID (a dropped or skipped one)
package resolve

import (
	"context"
	"fmt"
)

// retiredID checks snapshot b, the next version of the document at k, against
// the newest stored version (latestStored, rebuilt by snapshotOf) and the
// list's NEXT-LOCAL-ID (retiredIn). A diff (DELTA-BASE) and a list without
// NEXT-LOCAL-ID are not checked.
func (d Dir) retiredID(k DocKey, lg versionLog, b []byte) error {
	if _, isDiff := deltaBaseOf(b); isDiff || lg.Next == 0 {
		return nil
	}
	var prev []byte
	if e, ok := latestStored(lg.Entries); ok {
		var err error
		if prev, err = d.snapshotOf(context.Background(), k, lg.Entries, e, 0); err != nil {
			return err
		}
	}
	if id, ok := retiredIn(prev, b, lg.Next); ok {
		return fmt.Errorf("geul/resolve: LOCAL ID #%d is not in the previous version but below NEXT-LOCAL-ID %d: a dropped ID is not used again and a new statement takes its ID from NEXT-LOCAL-ID", id, lg.Next)
	}
	return nil
}
