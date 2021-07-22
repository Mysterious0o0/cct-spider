package store

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/filter"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/findmap"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/splitsql"

	"github.com/xiaogogonuo/cct-spider/internal/pkg/subString"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"strings"
	"time"
)

var (
	t                               = time.Now().Format("20060102")
	preamble, epilogue, oneQuoteSql = splitsql.GetInsertBaseSQLCode(&callback.SqlValues{}, "t_dmbe_policy_news_info")
)

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
		regionPat, regionMap := findmap.RegionRuntime()
		if r := findmap.FindOne(regionPat, mes.Summary); r != "" {
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
		v, l := splitsql.GetQuotesAndValues(sqlValues)

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
