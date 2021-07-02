package month

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/provincecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/response"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"github.com/xiaogogonuo/cct-spider/pkg/set"
)

// 地区经济指标-分省月度数据-居民消费价格指数(上年同月=100)

func prepareData(regionCode, indexCode string, diff [][]string) (data [][]string) {
	data = make([][]string, 0)
	for _, r := range diff {
		tv := response.TargetValue{}
		tv.AcctYear = r[0][:4]
		tv.AcctMonth = r[0][4:6]
		tv.TargetValue = r[1]
		tv.TargetCode = indexcode.StatInner[indexCode]
		tv.TargetGUID = md5.MD5(tv.TargetCode)
		tv.TargetName = indexcode.CodeName[indexCode]
		tv.SourceTargetCode = indexCode
		tv.RegionCode = regionCode
		tv.RegionName = provincecode.CodeProvince[regionCode]
		tv.ValueGUID = md5.MD5(tv.TargetCode + regionCode + r[0])  // TODO: unstable
		fmt.Println(tv)
		data = append(data, tv.Row())
	}
	return
}

func Run() {
	for _, v := range provincecode.ProvinceCode {
		url := urllib.Param{
			M:              "QueryData",
			DBCode:         "fsyd",
			RowCode:        "zb",
			ColCode:        "sj",
			WdsWdCode:      "reg",
			WdsWdValueCode: v,
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Months(indexcode.CPI3StartYear, 1),
		}
		querySQL := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s' AND REGION_CODE = '%s'`
		queryRow := mysql.Query(fmt.Sprintf(querySQL, indexcode.CPI3Code, v))
		row := executor.Executor(url, typecode.ProvinceMonthDataCode, indexcode.CPI3Code)
		diff, err := set.Set{Src: queryRow}.Diff(row)
		if err != nil {
			continue
		}
		data := prepareData(v, indexcode.CPI3Code, diff)
		if len(data) == 0 {
			continue
		}
		tranSQL := mysql.Generator(data)
		mysql.Transaction(tranSQL)
	}
}
