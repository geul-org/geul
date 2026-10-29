//ff:func feature=pack type=helper control=sequence
//ff:what corrupt — an ErrCorrupt error naming what is broken
package pack

import "fmt"

// corrupt wraps ErrCorrupt with what failed.
func corrupt(what string) error {
	return fmt.Errorf("%w: %s", ErrCorrupt, what)
}
