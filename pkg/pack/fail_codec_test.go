//ff:type feature=pack type=model
//ff:what failCodec — a Codec whose Compress always fails, for writer error paths
package pack

// failCodec fails every Compress.
type failCodec struct{}
