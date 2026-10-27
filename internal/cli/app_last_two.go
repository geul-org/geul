//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.lastTwo — the newest live version of a document and the live one before it
package cli

import "errors"

// lastTwo reads the two highest versions of addr's document that are not
// deleted (resolve.Dir.Log, At).
func (a *app) lastTwo(addr string) ([]byte, []byte, error) {
	k, err := a.key(addr)
	if err != nil {
		return nil, nil, err
	}
	d, err := a.dir()
	if err != nil {
		return nil, nil, err
	}
	h, err := d.Log(k.Issuer, k.Doc)
	if err != nil {
		return nil, nil, err
	}
	var live []uint32
	for _, e := range h.Entries {
		if !e.Deleted {
			live = append(live, e.Version)
		}
	}
	if len(live) < 2 {
		return nil, nil, errors.New("diff: fewer than two versions; give two")
	}
	b, err := d.At(k.Issuer, k.Doc, live[len(live)-2])
	if err != nil {
		return nil, nil, err
	}
	n, err := d.At(k.Issuer, k.Doc, live[len(live)-1])
	return b, n, err
}
