//ff:func feature=pack type=writer control=sequence
//ff:what writer.verify — open the new manifest and run the path checks and Check before latest.txt moves
package pack

import "context"

// verify runs BuildOptions.Verify and Check against manifest name.
func (w *writer) verify(name string) error {
	if !w.o.Verify && w.o.Check == nil {
		return nil
	}
	r, err := Open(w.ctx, FileStore{Root: w.out}, Options{Manifest: name})
	if err == nil && w.o.Verify {
		err = w.verifyPaths(r)
	}
	if err == nil && w.o.Check != nil {
		err = w.o.Check(context.WithoutCancel(w.ctx), r)
	}
	return err
}
