package set

import (
	"fmt"
	"testing"
)

func TestSet_Diff(t *testing.T) {
	s := Set{src: [][]string{
		{"1990", "1"},
		{"1991", "2"},
	}}
	diff, err := s.Diff([][]string{
		{"1989", "0"},
		{"1990", "1"},
		{"1991", "2"},
	})
	fmt.Println(diff, err)
}

func BenchmarkSet_Diff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := Set{src: [][]string{
			{"1990", "1"},
			{"1991", "2"},
		}}
		_, _ = s.Diff([][]string{
			{"1989", "0"},
			{"1990", "1"},
			{"1991", "2"},
		})
	}
}
