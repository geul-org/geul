//ff:func feature=cli type=helper control=sequence
//ff:what checkFile — the schema Facts of one file as findings
package cli

import (
	"os"
)

// checkFile reads path and checks its bytes (checkBytes).
func checkFile(path string) ([]finding, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return checkBytes(path, b), nil
}
