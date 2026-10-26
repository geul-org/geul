//ff:func feature=cli type=helper control=sequence
//ff:what app.readPair — two documents by file path or address
package cli

// readPair reads base and next (readDoc).
func (a *app) readPair(base, next string) ([]byte, []byte, error) {
	b, err := a.readDoc(base)
	if err != nil {
		return nil, nil, err
	}
	n, err := a.readDoc(next)
	return b, n, err
}
