package store

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/filter"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/subString"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	s                               = &callback.SqlValues{}
	t                               = time.Now().Format("20060102")
	preamble, epilogue, oneQuoteSql = _getSQL(s)
)

func InsertIntoSQL(f *filter.Filter, message <-chan *callback.Message) {
	var (
		quotes       []string
		insertValues []interface{}
		beginLen     = len(preamble) + len(epilogue)
	)
	for mes := range message {
		if len(mes.Title) == 0 && len(mes.Content) == 0 {
			continue
		}
		if mes.Date > t || mes.Date == "" {
			mes.Date = "10000101"
		}
		if len(mes.Content) > 65535 {
			n, _ := subString.RuneIndex([]byte(mes.Content), 65535/3)
			mes.Content = mes.Content[:n]
		}
		sqlValues := &callback.SqlValues{
			NEWS_GUID:        md5.MD5(mes.Date + mes.Title),
			NEWS_TITLE:       mes.Title,
			NEWS_TS:          mes.Date,
			NEWS_URL:         mes.Url,
			NEWS_SOURCE:      mes.Source,
			NEWS_SUMMARY:     mes.Content,
			POLICY_TYPE:      "10",
			POLICY_TYPE_NAME: "国家政策",
			IS_CONTROL:       "否",
			IS_INVEST:        "否",
			IS_DEPOSIT:       "否",
			IS_FUND:          "否",
			IS_STOCK:         "否",
			IS_FINANCE:       "否",
			IS_INDUSTRY:      "否",
			IS_CAPITAL:       "否",
		}
		f.WriteMap(md5.MD5(mes.Url))
		v, l := _getQuotesAndValues(sqlValues)

		if beginLen+l+len(oneQuoteSql) < 500000 {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			beginLen += len(oneQuoteSql) + l

		} else {
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			mysql.Transaction(SQl, insertValues...)
			f.SaveUrlKey()
			insertValues = append([]interface{}{}, v...)
			quotes = append([]string{}, oneQuoteSql)
			beginLen = len(preamble) + len(epilogue) + len(oneQuoteSql) + l
		}
	}
	SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
	mysql.Transaction(SQl, insertValues...)
	f.SaveUrlKey()
}

func _getQuotesAndValues(v interface{}) (insertValues []interface{}, strLen int) {
	insertValues = make([]interface{}, 0)

	elem := reflect.ValueOf(v)

	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	for i := 0; i < elem.NumField(); i++ {
		curField := elem.Field(i)
		val, valLen := _parseField(curField)
		strLen += valLen
		insertValues = append(insertValues, val)
	}
	return
}

func _parseField(v reflect.Value) (fieldValue interface{}, valLen int) {

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fieldValue = v.Int()
		valLen = len(strconv.FormatInt(v.Int(), 10))

	case reflect.String:
		fieldValue = v.String()
		valLen = v.Len()
	}
	return

}

func _getSQL(v interface{}) (preambleSql string, epilogueSql string, oneQuoteSql string) {

	elem := reflect.ValueOf(v)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}

	elemType := elem.Type()
	numFields := elem.NumField()

	quotes := strings.Repeat("?,", numFields)
	quotes = quotes[0 : len(quotes)-1]

	insertFields := make([]string, 0, numFields)
	epilogues := make([]string, 0, numFields)

	for i := 0; i < numFields; i++ {
		field := elemType.Field(i).Name
		insertFields = append(insertFields, field)
		epilogues = append(epilogues, fmt.Sprintf("%s = VALUES(%s)", field, field))

	}
	oneQuoteSql = fmt.Sprintf("(%s)", quotes)
	preambleSql = fmt.Sprintf("INSERT INTO t_dmbe_policy_news_info (%s) VALUES ", strings.Join(insertFields, ", "))
	epilogueSql = fmt.Sprintf("ON DUPLICATE KEY UPDATE %s", strings.Join(epilogues, ", "))
	return
}
