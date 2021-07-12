package core

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/indexcode"
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
	IndexCode   string
	TypeCode    string
	UnitType    string
	UnitName    string
	URL         urllib.Param
}

func (c Core) Run() (rowsAffected bool, err error) {
	queryRow := mysql.Query(c.SQL)
	row := executor.Executor(c.URL, c.TypeCode, c.IndexCode)
	diff, err := set.Set{Src: queryRow}.Diff(row)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	dataBuilder := response.DataBuilder{
		IndexCode: c.IndexCode,
		RegionCode: c.URL.WdsWdValueCode,
		TL: c.TL,
		UnitType: c.UnitType,
		UnitName: c.UnitName,
	}
	data := dataBuilder.Build(diff)
	if len(data) == 0 {
		logger.Info(fmt.Sprintf("%s has no new data to update",indexcode.CodeName[c.IndexCode]))
		return
	}
	tranSQL := mysql.Generator(data)
	mysql.Transaction(tranSQL)
	rowsAffected = true
	return
}