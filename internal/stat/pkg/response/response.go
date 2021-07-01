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
	TargetCODE       string // 指标编码
	TargetNAME       string // 指标名称
	SourceTargetCODE string // 来源系统指标编号
	RegionCODE       string // 统计地区
	RegionNAME       string // 统计地区说明
	UnitTYPE         string // 计量单位类型
	UnitNAME         string // 计量单位名称
	AcctYEAR         string // 年
	AcctQUARTOR      string // 季
	AcctMONTH        string // 月
	AcctDate         string // 日
	TargetVALUE      string // 指标值
}
