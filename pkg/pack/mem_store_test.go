//ff:type feature=pack type=model
//ff:what memStore — an in-memory Store for crafted objects
package pack

// memStore maps object names to bytes.
type memStore map[string][]byte
