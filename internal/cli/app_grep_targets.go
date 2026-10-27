//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.grepTargets — the documents to search: the arguments, or every document of the data folder by address
package cli

// grepTargets returns args, or every `<Issuer>/<문서>` of the data folder
// when there are none.
func (a *app) grepTargets(args []string) ([]string, error) {
	if len(args) > 0 {
		return args, nil
	}
	d, err := a.dir()
	if err != nil {
		return nil, err
	}
	issuers, err := d.Issuers()
	if err != nil {
		return nil, err
	}
	var out []string
	for _, is := range issuers {
		docs, err := a.documentsOf(d, is.ID)
		if err != nil {
			return nil, err
		}
		out = append(out, addressesOf(is.ID, docs)...)
	}
	return out, nil
}
