//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.ls — print Issuers or one Issuer's documents: ID and names, main name last
package cli

import (
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// ls lists the data folder's Issuers (resolve.Dir.Issuers), or with an Issuer
// (ID, name or local key name) its documents (resolve.Dir.Documents).
func (a *app) ls(args []string) error {
	d, err := a.dir()
	if err != nil {
		return err
	}
	var list []resolve.Named
	if len(args) == 0 {
		list, err = d.Issuers()
	} else {
		list, err = a.documentsOf(d, args[0])
	}
	if err != nil {
		return err
	}
	if done, err := a.json(list); done {
		return err
	}
	for _, n := range list {
		fmt.Fprintln(a.out, strings.TrimSpace(n.ID+" "+strings.Join(n.Names, " ")))
	}
	return nil
}
