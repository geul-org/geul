//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what appendColumns — each column's values as varints, column after column
package pack

import "encoding/binary"

// appendColumns writes the columns in order, each as len varints.
func appendColumns(out []byte, cols ...[]uint64) []byte {
	for _, col := range cols {
		for _, v := range col {
			out = binary.AppendUvarint(out, v)
		}
	}
	return out
}
