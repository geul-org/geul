//ff:func feature=pack type=helper control=sequence
//ff:what sortKVs — entries in ascending key order
package pack

import "sort"

// sortKVs sorts kvs by key and returns it.
func sortKVs(kvs []kv) []kv {
	sort.Slice(kvs, func(i, j int) bool { return kvs[i].key < kvs[j].key })
	return kvs
}
