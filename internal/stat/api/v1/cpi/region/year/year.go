package year

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

// 地区经济指标-分省年度数据-居民消费价格指数(上年=100)

func prepareData(regionCode, indexCode string, diff [][]string) (data [][]string) {
	data = make([][]string, 0)
	for _, r := range diff {
		tv := response.TargetValue{}
		tv.AcctYear = r[0]
		tv.TargetValue = r[1]
		tv.TargetCode = indexcode.StatInner[indexCode]
		tv.TargetGUID = md5.MD5(tv.TargetCode)
		tv.TargetName = indexcode.CodeName[indexCode]
		tv.SourceTargetCode = indexCode
		tv.RegionCode = regionCode
		tv.RegionName = provincecode.CodeProvince[regionCode]
		tv.ValueGUID = md5.MD5(tv.TargetCode + regionCode + tv.AcctYear)  // TODO: unstable
		fmt.Println(tv)
		data = append(data, tv.Row())
	}
	return
}

func Run() {
	for _, v := range provincecode.ProvinceCode {
		url := urllib.Param{
			M:              "QueryData",
			DBCode:         "fsnd",
			RowCode:        "zb",
			ColCode:        "sj",
			WdsWdCode:      "reg",
			WdsWdValueCode: v,  // 省、市编码
			DfWdsWdCode:    "sj",
			DfWdsValueCode: last.Years(indexcode.CPIStartYear),
		}
		querySQL := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s' AND REGION_CODE = '%s'`
		queryRow := mysql.Query(fmt.Sprintf(querySQL, indexcode.CPICode, v))
		row := executor.Executor(url, typecode.ProvinceYearDataCode, indexcode.CPICode)
		diff, err := set.Set{Src: queryRow}.Diff(row)
		if err != nil {
			continue
		}
		data := prepareData(v, indexcode.CPICode, diff)
		if len(data) == 0 {
			continue
		}
		tranSQL := mysql.Generator(data)
		mysql.Transaction(tranSQL)
	}
}