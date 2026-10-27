//ff:type feature=schema type=model
//ff:what Options — Strict turns Warns into Errors; Codecs adds layer-2 codec checks
package schema

// Options configures Check. Strict treats reserved codes as errors. Codecs are
// layer-2 checks run on every record they match; empty means layer 1 only.
type Options struct {
	Strict bool
	Codecs []Codec
}
