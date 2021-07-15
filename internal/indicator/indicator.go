package indicator

const (
	Annual    uint8 = iota // 年度
	Quarterly              // 季度
	Monthly                // 月度
	Weekly                 // 周度
	Daily                  // 日度
)

type Unit struct {
	UnitCode      string // 指标单位代码
	UnitName      string // 指标单位名称
}

type Region struct {
	RegionCode    string // 指标所属区域代码
	RegionName    string // 指标所属区域名称
}

type StartDate struct {
	StartYear     uint16 // 指标起始年
	StartQuarter  uint8  // 指标起始季
	StartMonth    uint8  // 指标起始月
	StartWeek     uint8  // 指标起始周
	StartDay      uint8  // 指标起始日
}

type Indicator struct {
	Name          string // 指标名称
	ShortName     string // 指标名简写
	ChenTongCode  string // 中国诚通指标代码
	StatisticCode string // 国家统计局指标代码
	Division      uint8  // 指标是年度、季度、月度、周度、日度
	Unit
	Region
	StartDate
}
