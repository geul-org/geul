//ff:func feature=delta type=encoder control=iteration dimension=1
//ff:what Diff — the diff that turns base into next: next's HEAD with DELTA-BASE and DELTA-REMOVE, the changed and new Edges; an error when Apply would not give next back
package delta

import (
	"bytes"
	"errors"
	"slices"

	"github.com/geul-org/geul/go/pkg/core"
)

// Diff writes the diff of next against base, version baseVersion of the
// same document (meta/lineage/DELTA-BASE): next's HEAD plus DELTA-BASE and,
// when Edges were removed, DELTA-REMOVE; the BODY holds the Edges next adds
// or changes (by Edge Id). The diff must give next back through Apply byte
// for byte; when next reorders the base's Edges or puts new ones before old
// ones it cannot, and Diff fails (keep next as a snapshot).
func Diff(base, next []byte, baseVersion uint32) ([]byte, error) {
	b, err := splitSnapshot(base)
	if err != nil {
		return nil, err
	}
	n, err := splitSnapshot(next)
	if err != nil {
		return nil, err
	}
	old := map[uint64]core.Record{}
	for i, id := range b.ids {
		old[id] = b.body[i]
	}
	var body []core.Record
	for i, id := range n.ids {
		if o, ok := old[id]; !ok || !slices.Equal(o.Payload, n.body[i].Payload) || o.Codec != n.body[i].Codec {
			body = append(body, n.body[i])
		}
	}
	head := append(slices.Clone(n.head), core.EncodeDeltaBase(baseVersion))
	if gone := removedIDs(b.ids, n.ids); len(gone) > 0 {
		head = append(head, core.EncodeDeltaRemove(gone))
	}
	d := encodeDoc(head, body)
	if got, err := Apply(base, d); err != nil || !bytes.Equal(got, next) {
		return nil, errors.New("geul/delta: next does not keep the base order with new Edges at the end; keep it as a snapshot")
	}
	return d, nil
}
