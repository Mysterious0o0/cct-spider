package store

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"strings"
	"sync"
)

var (
	values    []string
	valuesLen = 0
	preamble  = "INSERT INTO ministries.laws (ID, TITLE, CONTEXT, SOURCE, FILE_DATE) VALUES"
	epilogue  = "ON DUPLICATE KEY UPDATE ID = VALUES(ID), SOURCE = VALUES(SOURCE), TITLE = VALUES(TITLE), CONTEXT = VALUES(CONTEXT), FILE_DATE = VALUES(FILE_DATE);"
)

func InsertIntoSQL(message <- chan *Message, wg *sync.WaitGroup) {
	defer wg.Done()
	for mes := range message {
		if len(mes.Title) == 0 && len(mes.Content) == 0 {
			continue
		}
		id := md5.MD5(mes.Date + mes.Title)
		mesLen := len(id) + len(mes.Title) + len(mes.Content) + len(mes.Source) + len(mes.Date)
		if valuesLen + len(preamble) + len(epilogue) + mesLen < 10000 {
			values = append(values, fmt.Sprintf("(%s, %s, %s, %s, %s)", id, mes.Title, mes.Content, mes.Source, mes.Date))
			valuesLen += mesLen

		} else {
			v := strings.Join(values, ",")
			sqlCode := strings.Join([]string{preamble, v, epilogue}, " ")
			mysql.Transaction(sqlCode)
			values = append([]string{}, fmt.Sprintf("('%s', '%s', '%s', '%s', '%s')", id, mes.Title, mes.Content, mes.Source, mes.Date))
			valuesLen = mesLen
		}
	}
}
