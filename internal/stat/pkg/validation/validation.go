package validation

import "sort"

type Data [][]string

func(d Data) Len() int {
	return len(d)
}

func(d Data) Less(i, j int) bool {
	return d[i][0] > d[j][0] // descend
}

func(d Data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// LightEqual validate spider data is fake or not.
// dst is true data load from database
// if half of d is not equal to dst, comes out false
func (d Data) LightEqual(dst Data) bool {
	sort.Sort(d)
	sort.Sort(dst)
	c := 0
	for _, dV := range d {
		for _, dstV := range dst {
			if dV[0] == dstV[0] {
				if dV[1] == dstV[1] {
					c++
				}
				break
			}
		}
	}
	if c > d.Len() / 2 {
		return true
	}
	return false
}
