package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type M struct {
	A io.Reader
	B map[string]string
}

func main() {
	m := M{}
	fmt.Println(m.A, m.A == nil)
	fmt.Println(m.B, m.B == nil)
	for k, v := range m.B {
		fmt.Println(k,v)
	}
	//req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/test", nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//params := map[string]string{"name": "陆佳伟", "age": "29"}
	//q := req.URL.Query()
	//
	//for key, val := range params {
	//	q.Add(key, val)
	//}
	//req.URL.RawQuery = q.Encode()
	//fmt.Println(req.URL)

	router := gin.Default()

	router.GET("/viper", func(context *gin.Context) {
		context.JSON(http.StatusOK, nil)
	})

	_ = router.Run()
}
