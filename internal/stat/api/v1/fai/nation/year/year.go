package year

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/typecode"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/core"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/last"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"time"
)

// 主要经济指标-全国年度数据-固定资产投资价格指数(上年=100)
// 主要经济指标-全国年度数据-固定资产投资价格指数(1990=100)

func runFaiNationYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	faiRegion := last.YearRegion(indexcode.FAIStartYear)
	for _, region := range faiRegion {
		c := core.Core{
			TL: "year",
			SQL: fmt.Sprintf(sql, indexcode.FAICode),
			IndexCode: indexcode.FAICode,
			TypeCode: typecode.YearDataCode,
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgnd",
				RowCode:        "zb",
				ColCode:        "sj",
				DfWdsWdCode:    "sj",
				DfWdsValueCode: region,
			},
		}
		rowsAffected, err := c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 10)
	}

}

func Run() {
	runFaiNationYear()
}