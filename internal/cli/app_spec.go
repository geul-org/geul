//ff:func feature=cli type=helper control=sequence
//ff:what app.spec — the codec specs published in the data folder; empty without a data folder or when they do not read
package cli

import "github.com/geul-org/gogeul/pkg/codecspec"

// spec gathers the codec spec documents of the data folder
// (resolve.Dir.CodecSpec). Records of undeclared codecs are shown and read
// generically.
func (a *app) spec() codecspec.Spec {
	d, err := a.dir()
	if err != nil {
		return codecspec.Spec{}
	}
	s, err := d.CodecSpec()
	if err != nil {
		return codecspec.Spec{}
	}
	return s
}
