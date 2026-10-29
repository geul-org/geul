//ff:type feature=resolve type=model
//ff:what mapStore — an in-memory Store (fstest.MapFS plus writes) for Dir over a non-directory store
package resolve

import "testing/fstest"

// mapStore keeps paths in a MapFS.
type mapStore struct {
	fstest.MapFS
}
