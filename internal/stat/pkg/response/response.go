package response

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
	TargetCode       string // 指标编码
	TargetName       string // 指标名称
	SourceTargetCode string // 来源系统指标编号
	RegionCode       string // 统计地区
	RegionName       string // 统计地区说明
	UnitType         string // 计量单位类型
	UnitName         string // 计量单位名称
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
		tv.SourceTargetCode,
		tv.RegionCode,
		tv.RegionName,
		tv.UnitType,
		tv.UnitName,
		tv.AcctYear,
		tv.AcctSeason,
		tv.AcctMonth,
		tv.AcctDate,
		tv.TargetValue,
	}
	return
}