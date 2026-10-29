//ff:func feature=pack type=helper control=sequence
//ff:what TestDefaultInt — below lo takes the default
package pack

import "testing"

func TestDefaultInt(t *testing.T) {
	if defaultInt(0, 1, 9) != 9 || defaultInt(3, 1, 9) != 3 {
		t.Fatal("defaultInt")
	}
}
