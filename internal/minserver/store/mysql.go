package store

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/filter"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/findregion"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/subString"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	regionPat                       string
	regionMap                       map[string][]string
	s                               = &callback.SqlValues{}
	t                               = time.Now().Format("20060102")
	preamble, epilogue, oneQuoteSql = _getSQL(s)
)

func init() {
	var regionReg []string
	regionMap = make(map[string][]string)
	sqlCode := "SELECT REGION_NAME, REGION_CODE FROM T_DMAA_REGIOIN_CODE"
	for _, region := range mysql.Query(sqlCode) {
		reg := regexp.MustCompile("省|市|自治区|自治州|特别行政区|区|盟")
		r := reg.ReplaceAllString(region[0], ")")
		regionReg = append(regionReg, r)
		k := strings.Replace(r, ")", "", -1)
		if _, ok := regionMap[k]; !ok {
			regionMap[k] = region
		}
	}
	regionPat = "(" + strings.Join(regionReg, "|(")
}

func InsertIntoSQL(f *filter.Filter, message <-chan *callback.Message) {
	var (
		region       string
		regionCode   string
		quotes       []string
		insertValues []interface{}
		beginLen     = len(preamble) + len(epilogue)
	)
	for mes := range message {
		if len(mes.Title) == 0 && len(mes.Summary) == 0 {
			continue
		}
		if mes.Date == "" || mes.Date > t {
			mes.Date = "10000101"
		}
		if len(mes.Summary) > 65535 {
			n, _ := subString.RuneIndex([]byte(mes.Summary), 65535/3)
			mes.Summary = mes.Summary[:n]
		}
		if r := findregion.FindRegion(regionPat, mes.Summary); r != "" {
			region = regionMap[r][0]
			regionCode = regionMap[r][1]
		}
		sqlValues := &callback.SqlValues{
			NEWS_GUID:        md5.MD5(mes.Date + mes.Title),
			NEWS_TITLE:       mes.Title,
			NEWS_TS:          mes.Date,
			NEWS_URL:         mes.Url,
			NEWS_SOURCE:      mes.Source,
			NEWS_SOURCE_CODE: mes.SourceCode,
			NEWS_SUMMARY:     mes.Summary,
			POLICY_TYPE:      "10",
			POLICY_TYPE_NAME: "国家政策",
			REGION_NAME:      region,
			REGION_CODE:      regionCode,
			IS_CONTROL:       "N",
			IS_INVEST:        "N",
			IS_DEPOSIT:       "N",
			IS_FUND:          "N",
			IS_STOCK:         "N",
			IS_FINANCE:       "N",
			IS_INDUSTRY:      "N",
			IS_CAPITAL:       "N",
			NEWS_GYS_CODE:    "90",
			NEWS_GYS_NAME:    "爬虫",
		}
		//logger.Info("Success", logger.Field("url", mes.Url))
		f.WriteMap(md5.MD5(mes.Url))
		v, l := _getQuotesAndValues(sqlValues)

		//insertValues = append([]interface{}{}, v...)
		//quotes = append([]string{}, oneQuoteSql)
		//SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
		//mysql.Transaction(SQl, insertValues...)

		if beginLen+l+len(oneQuoteSql) < 500000 {
			insertValues = append(insertValues, v...)
			quotes = append(quotes, oneQuoteSql)
			beginLen += len(oneQuoteSql) + l

		} else {
			fmt.Printf("insert into data:%v , url: %v\n", len(insertValues), len(f.ThisUrlKey))
			SQl := fmt.Sprintf("%s%s %s", preamble, strings.Join(quotes, ", "), epilogue)
			mysql.Transaction(SQl, insertValues...)
			f.SaveUrlKey()
			insertValues = append([]interface{}{}, v...)
			quotes = append([]string{}, oneQuoteSql)
			beginLen = len(preamble) + len(epilogue) + len(oneQuoteSql) + l
		}
	}
	fmt.Printf("insert into data:%v , url: %v\n", len(insertValues), len(f.ThisUrlKey))
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
