//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.envValues — geul env KEY…: one value per line
package cli

import "fmt"

// envValues prints the value of each key, one per line (empty when unset).
func (a *app) envValues(keys []string) error {
	for _, k := range keys {
		fmt.Fprintln(a.out, a.cfg.Get(k))
	}
	return nil
}
