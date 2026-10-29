//ff:func feature=pack type=helper control=sequence
//ff:what BuildOptions.withDefaults — fill zero or unusable values with the recommended parameters (§6)
package pack

import "time"

// withDefaults returns o with defaults; fan-outs below 2 would never shrink a level.
func (o BuildOptions) withDefaults() BuildOptions {
	if o.PackSize <= 0 {
		o.PackSize = DefaultPackSize
	}
	o.PLeaf = defaultInt(o.PLeaf, 1, DefaultPLeaf)
	o.HLeaf = defaultInt(o.HLeaf, 1, DefaultHLeaf)
	o.HFanout = defaultInt(o.HFanout, 2, DefaultHFanout)
	o.HRoot = defaultInt(o.HRoot, 1, DefaultHRoot)
	o.HLayers = defaultInt(o.HLayers, 1, DefaultHLayers)
	if o.Codec == nil {
		o.Codec = Stored{}
	}
	if o.Now == nil {
		o.Now = time.Now
	}
	return o
}
