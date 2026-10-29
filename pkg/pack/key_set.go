//ff:type feature=pack type=model
//ff:what keySet — key → location where the first (newest) layer to report a key wins
package pack

// keySet collects layer entries newest layer first.
type keySet map[string]Loc
