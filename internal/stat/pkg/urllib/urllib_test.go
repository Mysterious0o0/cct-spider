package urllib

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	p := Param{
		M:              "QueryData",
		DBCode:         "fsyd",
		RowCode:        "zb",
		ColCode:        "sj",
		WdsWdCode:      "reg",
		WdsWdValueCode: "110000",
		DfWdsWdCode:    "sj",
		DfWdsValueCode: 10,
	}
	fmt.Println(p.Encode())
}
