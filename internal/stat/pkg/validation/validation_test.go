package validation

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := Data{
		{"1990", "1"},
		{"1989", "2"},
		{"1991", "0"},
		{"1888", "10"},
		{"2002", "11"},
	}
	sort.Sort(data)
	d2 := Data{
		{"1990", "10"},
		{"1989", "20"},
		{"1991", "00"},
		{"1888", "10"},
		{"2002", "11"},
	}
	o := data.LightEqual(d2)
	fmt.Println(o)
}
