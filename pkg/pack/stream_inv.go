//ff:type feature=pack type=model
//ff:what streamInv — invEntry (§4.5) with exported fields for a checkpoint
package pack

// streamInv mirrors invEntry (§4.5) so a checkpoint can hold it (gob).
type streamInv struct {
	Kind  byte
	Enc   byte
	Flags uint16
	Ref   uint32
	Doc   uint64
	Slot  uint32
	N     uint32
	Off   uint64
}
