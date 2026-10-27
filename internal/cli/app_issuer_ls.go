//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.issuerLs — print each local Issuer's name and ID
package cli

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// issuerLs lists localKeys as `name ID`.
func (a *app) issuerLs() error {
	keys, err := a.localKeys()
	if err != nil {
		return err
	}
	for _, k := range keys {
		fmt.Fprintf(a.out, "%s %s\n", k.Name, core.FormatID(k.Issuer))
	}
	return nil
}
