package search

func BinarySearch(a []int, x int) bool {
	l := 0
	u := len(a)-1

	for u >= l {
		m := (u-l)/2+l
		if a[m] == x {
			return true
		} else if a[m] > x {
			u = m-1
		} else {
			l = m+1
		}
	}
	return false
}
