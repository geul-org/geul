//ff:func feature=cli type=command control=selection
//ff:what app.publish — check a draft and publish it (or amend), then optionally name it; print the version and hash
package cli

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// publish refuses a draft with schema errors (schemaErrors), stores it with
// resolve.Dir.Publish, or Overwrite for amend, binds name when given
// (resolve.Dir.Name) and prints `<버전> 0x<내용 해시>`. Content the document
// already has as a version is not published again (edge/PUBLICATION): that
// version is printed with a note, even when it is not the newest.
func (a *app) publish(path string, amend bool, name string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := schemaErrors(b); err != nil {
		return err
	}
	d, err := a.dir()
	if err != nil {
		return err
	}
	seen := hashesOf(d, b)
	var e resolve.VersionEntry
	switch {
	case amend:
		e, err = d.Overwrite(b)
	default:
		e, err = d.Publish(b)
	}
	if err != nil {
		return err
	}
	note := ""
	if !amend && seen[e.Hash] {
		note = " (already published; nothing written)"
	}
	fmt.Fprintf(a.out, "%d 0x%s%s\n", e.Version, hex.EncodeToString(e.Hash[:]), note)
	if name == "" {
		return nil
	}
	issuer, doc, _ := docIDOf(b)
	return d.Name(issuer, doc, name)
}
