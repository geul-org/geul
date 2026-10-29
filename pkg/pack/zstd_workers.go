//ff:func feature=pack type=helper control=sequence
//ff:what Zstd.workers — how many records compress at once (0 is DefaultWorkers)
package pack

// workers is z.Workers or DefaultWorkers.
func (z *Zstd) workers() int {
	return defaultInt(z.Workers, 1, DefaultWorkers)
}
