package main

import (
	"fmt"
	M "github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
)

func main() {
	r := M.Query("SELECT * FROM t_dmaa_base_target")
	fmt.Println(len(r))
}
