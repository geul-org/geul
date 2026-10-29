//ff:func feature=pack type=writer control=sequence
//ff:what StreamOptions.withDefaults — block and pack sizes as for Write, Zstd when no codec, the default run size
package pack

// withDefaults fills o.Build as Write does, makes a nil codec Zstd and a run
// size below 1 DefaultRunEntries.
func (o StreamOptions) withDefaults() StreamOptions {
	o.Build = o.Build.withDefaults()
	if o.Build.Codec == nil {
		o.Build.Codec = &Zstd{}
	}
	o.RunEntries = defaultInt(o.RunEntries, 1, DefaultRunEntries)
	return o
}
