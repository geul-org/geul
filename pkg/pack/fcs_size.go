//ff:func feature=pack type=decoder control=sequence
//ff:what fcsSize — the Frame_Content_Size field length from its flag and Single_Segment_Flag
package pack

// fcsSize is 0, 1, 2, 4 or 8 (flag 0 means 1 byte only in a single-segment frame).
func fcsSize(flag byte, single bool) int {
	if flag == 0 && single {
		return 1
	}
	return [4]int{0, 2, 4, 8}[flag&3]
}
