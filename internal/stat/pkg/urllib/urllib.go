package urllib

import (
	"fmt"
	"net/url"
)

const (
	base = "https://data.stats.gov.cn/easyquery.htm?"
)

type Param struct {
	M              string
	DBCode         string
	RowCode        string
	ColCode        string
	WdsWdCode      string
	WdsWdValueCode string
	DfWdsWdCode    string
	DfWdsValueCode int
}

func (p Param) Encode() string {
	value := url.Values{}
	value.Set("m", p.M)
	value.Set("dbcode", p.DBCode)
	value.Set("rowcode", p.RowCode)
	value.Set("colcode", p.ColCode)
	if p.WdsWdCode != "" || p.WdsWdValueCode != "" {
		value.Set("wds", fmt.Sprintf(`[{"wdcode":"%s","valuecode":"%s"}]`, p.WdsWdCode, p.WdsWdValueCode))
	} else {
		value.Set("wds", `[]`)
	}
	value.Set("dfwds", fmt.Sprintf(`[{"wdcode":"%s","valuecode":"last%d"}]`, p.DfWdsWdCode, p.DfWdsValueCode))
	v := value.Encode()
	return base + v
}

//m: QueryData
//dbcode: fsnd  // 分省年度
//rowcode: zb
//colcode: sj
//wds: [{"wdcode":"reg","valuecode":"110000"}]
//dfwds: [{"wdcode":"sj","valuecode":"last25"}]

//m: QueryData
//dbcode: fsyd  // 分省月度
//rowcode: zb
//colcode: sj
//wds: [{"wdcode":"reg","valuecode":"110000"}]
//dfwds: [{"wdcode":"sj","valuecode":"last24"}]

// 指标年度
//m: QueryData
//dbcode: hgnd
//rowcode: zb
//colcode: sj
//wds: []
//dfwds: [{"wdcode":"sj","valuecode":"last10"}]

// 指标月度
//m: QueryData
//dbcode: hgyd
//rowcode: zb
//colcode: sj
//wds: []
//dfwds: [{"wdcode":"sj","valuecode":"last8"}]

// 指标季度
//m: QueryData
//dbcode: hgjd
//rowcode: zb
//colcode: sj
//wds: []
//dfwds: [{"wdcode":"sj","valuecode":"last100"}]