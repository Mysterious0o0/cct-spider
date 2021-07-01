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

func prepareData(v string, diff [][]string) (data [][]string) {
	data = make([][]string, 0)
	for _, r := range diff {
		tv := response.TargetValue{}
		tv.AcctYEAR = r[0]
		tv.TargetVALUE = r[1]
		tv.TargetCODE = indexcode.StatInner[indexcode.CPICode]
		tv.TargetGUID = md5.MD5(tv.TargetCODE)
		tv.TargetNAME = indexcode.CodeName[indexcode.CPICode]
		tv.SourceTargetCODE = indexcode.CPICode
		tv.RegionCODE = v
		tv.RegionNAME = provincecode.CodeProvince[v]
		tv.ValueGUID = md5.MD5(tv.TargetCODE + v + tv.AcctYEAR)
		fmt.Println(tv)
		data = append(data, []string{
			tv.ValueGUID,
			tv.TargetGUID,
			tv.TargetCODE,
			tv.TargetNAME,
			tv.SourceTargetCODE,
			tv.RegionCODE,
			tv.RegionNAME,
			tv.UnitTYPE,
			tv.UnitNAME,
			tv.AcctYEAR,
			tv.AcctQUARTOR,
			tv.AcctMONTH,
			tv.AcctDate,
			tv.TargetVALUE,
		})
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
		data := prepareData(v, diff)
		if len(data) == 0 {
			continue
		}
		tranSQL := mysql.Generator(data)
		mysql.Transaction(tranSQL)
	}
}