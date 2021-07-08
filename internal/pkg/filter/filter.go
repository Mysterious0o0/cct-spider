package filter

import (
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"io/ioutil"
	"os"
	"strings"
)

type Filter struct {
	Filepath string
	UrlKey   map[string]byte
}


func (f Filter)SaveUrlKey(content  []byte){
	w1, err := os.OpenFile(f.Filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil{
		logger.Error(err.Error())
		return
	}
	_, err = w1.Write(content)
	if err != nil{
		logger.Error(err.Error())
		return
	}

	err = w1.Close()
	if err != nil{
		logger.Error(err.Error())
		return
	}

}


func (f Filter)ReadUrlKey() {
	f.UrlKey = make(map[string]byte)
	fi, err := os.Open(f.Filepath)
	if err != nil{
		return
	}
	fd, err := ioutil.ReadAll(fi)
	if err != nil{
		logger.Error(err.Error())
		return
	}
	err = fi.Close()
	if err != nil{
		logger.Error(err.Error())
		return
	}
	for _, k := range strings.Split(string(fd), "\n"){
		if _, ok := f.UrlKey[k]; !ok{
			f.UrlKey[k] = 0
		}
	}
}
