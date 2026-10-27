//ff:func feature=cli type=command control=sequence
//ff:what geul CLI 엔트리포인트
package main

import (
	"os"

	"github.com/geul-org/gogeul/internal/cli"
)

func main() {
	os.Exit(cli.Execute())
}
