//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.show — print a document's HEAD lines, BODY, and BODY lines
package cli

import "fmt"

// show reads the document (readDoc) and prints its lines (linesOf), or them
// as JSON.
func (a *app) show(arg string) error {
	b, err := a.readDoc(arg)
	if err != nil {
		return err
	}
	ls, err := a.linesOf(b)
	if err != nil {
		return err
	}
	if done, err := a.json(ls); done {
		return err
	}
	for _, l := range ls.Head {
		fmt.Fprintln(a.out, l)
	}
	fmt.Fprintln(a.out, "BODY")
	for _, l := range ls.Body {
		fmt.Fprintln(a.out, l)
	}
	return nil
}
