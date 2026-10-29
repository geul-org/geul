//ff:func feature=cli type=helper control=sequence
//ff:what dictNote — ", dictionary ID" for a write that trained a zstd dictionary, "" otherwise
package cli

import "fmt"

// dictNote names a new dictionary for printResults.
func dictNote(id uint32) string {
	if id == 0 {
		return ""
	}
	return fmt.Sprintf(", dictionary %d", id)
}
