//ff:func feature=pack type=helper control=sequence
//ff:what notFound — resolve.ErrNotFound naming the path, so a resolver moves on to the next source
package pack

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// notFound reports that this generation has no bytes at path.
func notFound(path string) error {
	return fmt.Errorf("%w: %s", resolve.ErrNotFound, path)
}
