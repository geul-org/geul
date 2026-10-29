//ff:func feature=pack type=writer control=sequence
//ff:what writeTo — Write over a given fileSys: open the previous generation, plan, write data, index, manifest
package pack

import (
	"context"
	"path/filepath"
)

// writeTo is Write with the file system injected.
func writeTo(ctx context.Context, out string, b Batch, o BuildOptions, fs fileSys) (BuildResult, error) {
	w := &writer{ctx: ctx, out: out, o: o.withDefaults(), fs: fs, b: b}
	if err := w.open(); err != nil {
		return BuildResult{}, err
	}
	if err := w.plan(); err != nil {
		return BuildResult{}, err
	}
	if len(w.puts) == 0 && len(w.dels) == 0 {
		return w.res, nil
	}
	defer w.abort()
	err := fs.MkdirAll(filepath.Join(out, "o"))
	if err == nil {
		err = fs.MkdirAll(filepath.Join(out, "m"))
	}
	if err == nil {
		err = w.writeData()
	}
	if err == nil {
		err = w.finish()
	}
	if err != nil {
		return BuildResult{}, err
	}
	return w.res, nil
}
