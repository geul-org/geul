//ff:type feature=toulmin type=model
//ff:what CodecRef — where a TANGEUL record type lives: a direct Edge codec or a GEUL ISSUER format number
package toulmin

// CodecRef identifies a TANGEUL record type on the wire: either a direct TOULMIN
// Edge codec (Format == 0) or the ISSUER-1WORD codec 0x7C with a GEUL format
// number (Codec == 0x7C, Format != 0).
type CodecRef struct {
	Codec  uint8
	Format uint16
}

// codecIssuer1Word is the Edge ISSUER-1WORD codec used for every format number.
const codecIssuer1Word uint8 = 0x7C
