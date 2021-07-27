package response

import (
	"fmt"
	"github.com/xiaogogonuo/cct-spider/internal/stat/code/provincecode"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
)

type Response struct {
	ReturnData ReturnData `json:"returndata"`
}

type ReturnData struct {
	DataNodes []Node `json:"datanodes"`
}

type Node struct {
	Data Data
	Code string `json:"code"`
}

type Data struct {
	Data    float64 `json:"data"`
	StrData string  `json:"strdata"`
}

// TargetValue correspond table T_DMAA_BASE_TAREGT_VALUE
type TargetValue struct {
	ValueGUID        string // 指标值ID=md5(TargetCODE+RegionCODE+AcctYEAR+AcctQUARTOR+AcctMONTH+AcctDATE)
	TargetGUID       string // 指标ID
	TargetCode       string // 指标代码
	TargetName       string // 指标名称
	DataSourceCode   string // 数据来源代码
	DataSourceName   string // 数据来源名称
	SourceTargetCode string // 来源系统指标代码
	RegionCode       string // 统计地区
	RegionName       string // 统计地区说明
	IsQuantity       string // 是否定量
	UnitType         string // 计量单位类型
	UnitName         string // 计量单位名称
	PeriodType       string // 计量单位类型
	PeriodName       string // 计量单位名称
	AcctYear         string // 年
	AcctSeason       string // 季
	AcctMonth        string // 月
	AcctDate         string // 日
	TargetValue      string // 指标值
}

func (tv TargetValue) Row() (row []string) {
	row = []string{
		tv.ValueGUID,
		tv.TargetGUID,
		tv.TargetCode,
		tv.TargetName,
		tv.DataSourceCode,
		tv.DataSourceName,
		tv.SourceTargetCode,
		tv.RegionCode,
		tv.RegionName,
		tv.IsQuantity,
		tv.UnitType,
		tv.UnitName,
		tv.PeriodType,
		tv.PeriodName,
		tv.AcctYear,
		tv.AcctSeason,
		tv.AcctMonth,
		tv.AcctDate,
		tv.TargetValue,
	}
	return
}

type DataBuilder struct {
	TL         string // 时间类型：年度、季度、月度、日度(year、season、month、day)
	IndexName  string // 指标名称
	RegionCode string // 地区代码
	UnitType   string
	UnitName   string
	IndexMap   map[string]map[string]string
}

func (d DataBuilder) Build(diff [][]string) (data [][]string) {
	data = make([][]string, 0)
	for _, r := range diff {
		tv := TargetValue{}
		switch d.TL {
		case "year":
			tv.AcctYear = r[0]
		case "season":
			tv.AcctYear = r[0][:4]
			tv.AcctSeason = r[0][4:6]
		case "month":
			tv.AcctYear = r[0][:4]
			tv.AcctMonth = r[0][4:6]
		case "day":
			tv.AcctYear = r[0][:4]
			tv.AcctMonth = r[0][4:6]
			tv.AcctDate = r[0][6:8]
		}
		tv.TargetValue = r[1]
		tv.RegionCode = d.RegionCode
		tv.RegionName = provincecode.CodeProvince[d.RegionCode]
		tv.TargetCode = d.IndexMap[d.IndexName]["innerCode"]
		tv.TargetGUID = md5.MD5(tv.TargetCode)
		tv.TargetName = d.IndexName
		tv.SourceTargetCode = d.IndexMap[d.IndexName]["outerCode"]
		tv.ValueGUID = md5.MD5(tv.TargetCode + d.RegionCode + r[0])
		tv.DataSourceCode = "stat"
		tv.DataSourceName = "国家统计局"
		tv.UnitType = d.UnitType
		tv.UnitName = d.UnitName
		data = append(data, tv.Row())
	}
	logger.Info(fmt.Sprintf("%s update %d rows", d.IndexName, len(data)))
	return
}