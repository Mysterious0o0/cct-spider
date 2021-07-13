package core

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/executor"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/response"
	"github.com/xiaogogonuo/cct-spider/internal/stat/pkg/urllib"
	"github.com/xiaogogonuo/cct-spider/pkg/db/mysql"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"github.com/xiaogogonuo/cct-spider/pkg/set"
)

type Core struct {
	TL          string // 时间类型：年度、季度、月度、日度(year、season、month、day)
	SQL         string
	IndexName   string
	TypeCode    string
	UnitType    string
	UnitName    string
	URL         urllib.Param
	IndexMap    map[string]map[string]string
}

func (c Core) Run() (rowsAffected bool, err error) {
	indexOuterCode := c.IndexMap[c.IndexName]["outerCode"]
	queryRow := mysql.Query(c.SQL)
	row := executor.Executor(c.URL, c.TypeCode, indexOuterCode)
	diff, err := set.Set{Src: queryRow}.Diff(row)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	dataBuilder := response.DataBuilder{
		TL: c.TL,
		IndexName: c.IndexName,
		RegionCode: c.URL.WdsWdValueCode,
		UnitType: c.UnitType,
		UnitName: c.UnitName,
		IndexMap: c.IndexMap,
	}
	data := dataBuilder.Build(diff)
	if len(data) == 0 {
		logger.Info(fmt.Sprintf("%s no new data to update", c.IndexName))
		return
	}
	tranSQL := mysql.Generator(data)
	mysql.Transaction(tranSQL)
	rowsAffected = true
	return
}