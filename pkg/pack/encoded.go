//ff:type feature=pack type=model
//ff:what encoded — a record's stored form: encoding and bytes
package pack

// encoded is what encodeRecord made of one record.
type encoded struct {
	enc byte
	b   []byte
}
