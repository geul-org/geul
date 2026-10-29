//ff:func feature=cli type=helper control=sequence
//ff:what app.writeOptions — the BuildOptions of a pack write: the --compress value, else GEULCOMPRESS, as the codec (empty keeps each set's)
package cli

import "github.com/geul-org/gogeul/pkg/pack"

// writeOptions parses the codec setting; a flag value wins over GEULCOMPRESS.
func (a *app) writeOptions(flag string) (pack.BuildOptions, error) {
	s := flag
	if s == "" {
		s = a.cfg.Get("GEULCOMPRESS")
	}
	c, err := pack.ParseCodec(s)
	if err != nil {
		return pack.BuildOptions{}, err
	}
	o := packOptions()
	o.Codec = c
	return o, nil
}
