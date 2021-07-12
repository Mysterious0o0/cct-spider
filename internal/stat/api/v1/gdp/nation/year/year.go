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

// 主要经济指标-全国年度数据-国内生产总值年度(亿元)

func runGDPNationYear() {
	sql := `SELECT ACCT_YEAR, TARGET_VALUE FROM T_DMAA_BASE_TARGET_VALUE 
                WHERE SOURCE_TARGET_CODE = '%s'`

	gdpRegion := last.YearRegion(indexcode.GDPStartYear)
	for _, region := range gdpRegion {
		c := core.Core{
			TL: "year",
			SQL: fmt.Sprintf(sql, indexcode.GDPCode),
			IndexCode: indexcode.GDPCode,
			TypeCode: typecode.YearDataCode,
			UnitType: "",
			UnitName: "亿元",
			URL: urllib.Param{
				M:              "QueryData",
				DBCode:         "hgnd",
				RowCode:        "zb",
				ColCode:        "sj",
				DfWdsWdCode:    "sj",
				DfWdsValueCode: region,
			},
		}
		rowsAffected, err :=c.Run()
		if err != nil || !rowsAffected {
			return
		}
		time.Sleep(time.Second * 5)
	}
}

func Run() {
	runGDPNationYear()
}