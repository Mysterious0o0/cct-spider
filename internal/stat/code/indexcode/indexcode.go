package indexcode

// 指标名称(国家统计局和系统内部统一)
const (
	CPIName  = "居民消费价格指数(上年=100)"
	CPI1Name = "居民消费价格指数(1978=100)"
	CPI2Name = "居民消费价格指数(上月=100)"
	CPI3Name = "居民消费价格指数(上年同月=100)"
	CPI4Name = "居民消费价格指数(上年同期=100)"
	CQCName  = "货币和准货币(M2)供应量同比增长率"
	CQC1Name = "货币和准货币(M2)供应量_期末值"
	CQC2Name = "货币和准货币(M2)供应量_同比增长"
	FAIName  = "固定资产投资价格指数(上年=100)"
	FAI1Name = "固定资产投资价格指数(1990=100)"
	FAI2Name = "固定资产投资价格指数_当季值(上年同季=100)"
	FAI3Name = "固定资产投资价格指数_累计值(上年同期=100)"
	GDPName  = "国内生产总值"
	GDPRName = "地区生产总值"
	GDP1Name = "国内生产总值_当季值"
	GDP2Name = "国内生产总值_累计值"
	HCEName  = "居民人均消费支出"
	HCE1Name = "居民人均消费支出_同比增长"
	HCE2Name = "居民人均消费支出_累计值"
	HCE3Name = "居民人均消费支出_累计增长"
	IAVName  = "工业增加值"
	IAV1Name = "工业增加值_同比增长"
	IAV2Name = "工业增加值_累计增长"
	IAV3Name = "工业增加值_当季值"
	IAV4Name = "工业增加值_累计值"
	PMIName  = "制造业采购经理指数"
	PPIName  = "工业生产者出厂价格指数(上年=100)"
	PPI1Name = "工业生产者出厂价格指数(1985=100)"
	PPI2Name = "工业生产者出厂价格指数(上月=100)"
	PPI3Name = "工业生产者出厂价格指数(上年同月=100)"
	PPI4Name = "工业生产者出厂价格指数(上年同期=100)"
	RCLName  = "居民消费水平"
	SCGName  = "社会消费品零售总额"
	SCG1Name = "社会消费品零售总额_当期值"
	SCG2Name = "社会消费品零售总额_累计值"
	SCG3Name = "社会消费品零售总额_同比增长"
	SCG4Name = "社会消费品零售总额_累计增长"
	URIName  = "城镇居民人均可支配收入"
	URI1Name = "城镇居民人均可支配收入_同比增长"
	URI2Name = "城镇居民人均可支配收入_累计值"
	URI3Name = "城镇居民人均可支配收入_累计增长"
)

// IndexMap record index name mapping outer code and inner code
var IndexMap map[string]map[string]string

func setUpIndexMap() {
	IndexMap = make(map[string]map[string]string)

	// 居民消费价格指数(上年=100)
	IndexMap[CPIName] = make(map[string]string)
	IndexMap[CPIName]["outerCode"] = "A090101"
	IndexMap[CPIName]["innerCode"] = "HG00003"
	IndexMap[CPIName]["startYear"] = "1951"

	// 居民消费价格指数(1978=100)
	IndexMap[CPI1Name] = make(map[string]string)
	IndexMap[CPI1Name]["outerCode"] = "A090201"
	IndexMap[CPI1Name]["innerCode"] = "" // TODO: no need yet
	IndexMap[CPI1Name]["startYear"] = "1978"

	// 居民消费价格指数(上月=100)
	IndexMap[CPI2Name] = make(map[string]string)
	IndexMap[CPI2Name]["outerCode"] = "A01030101"
	IndexMap[CPI2Name]["innerCode"] = "HG00037"
	IndexMap[CPI2Name]["startYear"] = "2016"

	// 居民消费价格指数(上年同月=100)
	IndexMap[CPI3Name] = make(map[string]string)
	IndexMap[CPI3Name]["outerCode"] = "A01010101"
	IndexMap[CPI3Name]["innerCode"] = "HG00004"
	IndexMap[CPI3Name]["startYear"] = "2016"

	// 居民消费价格指数(上年同期=100)
	IndexMap[CPI4Name] = make(map[string]string)
	IndexMap[CPI4Name]["outerCode"] = "A01020101"
	IndexMap[CPI4Name]["innerCode"] = "HG00038"
	IndexMap[CPI4Name]["startYear"] = "2016"

	// 货币和准货币(M2)供应量同比增长率
	IndexMap[CQCName] = make(map[string]string)
	IndexMap[CQCName]["outerCode"] = "A0L0309"
	IndexMap[CQCName]["innerCode"] = "HG00005"
	IndexMap[CQCName]["startYear"] = "1991"

	// 货币和准货币(M2)供应量_期末值(亿元)
	IndexMap[CQC1Name] = make(map[string]string)
	IndexMap[CQC1Name]["outerCode"] = "A0D0101"
	IndexMap[CQC1Name]["innerCode"] = "HG00006"
	IndexMap[CQC1Name]["startYear"] = "1999"

	// 货币和准货币(M2)供应量_同比增长
	IndexMap[CQC2Name] = make(map[string]string)
	IndexMap[CQC2Name]["outerCode"] = "A0D0102"
	IndexMap[CQC2Name]["innerCode"] = "HG00007"
	IndexMap[CQC2Name]["startYear"] = "1999"

	// 固定资产投资价格指数(上年=100)
	IndexMap[FAIName] = make(map[string]string)
	IndexMap[FAIName]["outerCode"] = "A090107"
	IndexMap[FAIName]["innerCode"] = "HG00008"
	IndexMap[FAIName]["startYear"] = "1990"

	// 固定资产投资价格指数(1990=100)
	IndexMap[FAI1Name] = make(map[string]string)
	IndexMap[FAI1Name]["outerCode"] = "A090207"
	IndexMap[FAI1Name]["innerCode"] = "" // TODO: no need yet
	IndexMap[FAI1Name]["startYear"] = "1990"

	// 固定资产投资价格指数_当季值(上年同季=100)
	IndexMap[FAI2Name] = make(map[string]string)
	IndexMap[FAI2Name]["outerCode"] = "A060201"
	IndexMap[FAI2Name]["innerCode"] = "HG00009"
	IndexMap[FAI2Name]["startYear"] = "1998"

	// 固定资产投资价格指数_累计值(上年同期=100)
	IndexMap[FAI3Name] = make(map[string]string)
	IndexMap[FAI3Name]["outerCode"] = "A060301"
	IndexMap[FAI3Name]["innerCode"] = "HG00010"
	IndexMap[FAI3Name]["startYear"] = "2007"

	// 国内生产总值
	IndexMap[GDPName] = make(map[string]string)
	IndexMap[GDPName]["outerCode"] = "A020102"
	IndexMap[GDPName]["innerCode"] = "HG00001"
	IndexMap[GDPName]["startYear"] = "1952"

	// 地区生产总值
	IndexMap[GDPRName] = make(map[string]string)
	IndexMap[GDPRName]["outerCode"] = "A020101"
	IndexMap[GDPRName]["innerCode"] = "HG00002"
	IndexMap[GDPRName]["startYear"] = "1992"

	// 国内生产总值_当季值
	IndexMap[GDP1Name] = make(map[string]string)
	IndexMap[GDP1Name]["outerCode"] = "A010101"
	IndexMap[GDP1Name]["innerCode"] = "HG00035"
	IndexMap[GDP1Name]["startYear"] = "1992"

	// 国内生产总值_累计值
	IndexMap[GDP2Name] = make(map[string]string)
	IndexMap[GDP2Name]["outerCode"] = "A010102"
	IndexMap[GDP2Name]["innerCode"] = "HG00036"
	IndexMap[GDP2Name]["startYear"] = "1992"

	// 居民人均消费支出
	IndexMap[HCEName] = make(map[string]string)
	IndexMap[HCEName]["outerCode"] = "A0A0107"
	IndexMap[HCEName]["innerCode"] = "HG00011"
	IndexMap[HCEName]["startYear"] = "2013"

	// 居民人均消费支出_同比增长
	IndexMap[HCE1Name] = make(map[string]string)
	IndexMap[HCE1Name]["outerCode"] = "A0A0108"
	IndexMap[HCE1Name]["innerCode"] = "HG00012"
	IndexMap[HCE1Name]["startYear"] = "2014"

	// 居民人均消费支出_累计值
	IndexMap[HCE2Name] = make(map[string]string)
	IndexMap[HCE2Name]["outerCode"] = "A050109"
	IndexMap[HCE2Name]["innerCode"] = "HG00013"
	IndexMap[HCE2Name]["startYear"] = "2013"

	// 居民人均消费支出_累计增长
	IndexMap[HCE3Name] = make(map[string]string)
	IndexMap[HCE3Name]["outerCode"] = "A05010A"
	IndexMap[HCE3Name]["innerCode"] = "HG00014"
	IndexMap[HCE3Name]["startYear"] = "2014"

	// 工业增加值
	IndexMap[IAVName] = make(map[string]string)
	IndexMap[IAVName]["outerCode"] = "A020403"
	IndexMap[IAVName]["innerCode"] = "HG00015"
	IndexMap[IAVName]["startYear"] = "1952"

	// 工业增加值_同比增长
	IndexMap[IAV1Name] = make(map[string]string)
	IndexMap[IAV1Name]["outerCode"] = "A020101"
	IndexMap[IAV1Name]["innerCode"] = "HG00016"
	IndexMap[IAV1Name]["startYear"] = "1998"

	// 工业增加值_累计增长
	IndexMap[IAV2Name] = make(map[string]string)
	IndexMap[IAV2Name]["outerCode"] = "A020102"
	IndexMap[IAV2Name]["innerCode"] = "HG00017"
	IndexMap[IAV2Name]["startYear"] = "1998"

	// 工业增加值_当季值
	IndexMap[IAV3Name] = make(map[string]string)
	IndexMap[IAV3Name]["outerCode"] = "A01010B"
	IndexMap[IAV3Name]["innerCode"] = "HG00018"
	IndexMap[IAV3Name]["startYear"] = "1992"

	// 工业增加值_累计值
	IndexMap[IAV4Name] = make(map[string]string)
	IndexMap[IAV4Name]["outerCode"] = "A01010C"
	IndexMap[IAV4Name]["innerCode"] = "HG00019"
	IndexMap[IAV4Name]["startYear"] = "1992"

	// 制造业采购经理指数
	IndexMap[PMIName] = make(map[string]string)
	IndexMap[PMIName]["outerCode"] = "A0B0101"
	IndexMap[PMIName]["innerCode"] = "HG00020"
	IndexMap[PMIName]["startYear"] = "2005"

	// 工业生产者出厂价格指数(上年=100)
	IndexMap[PPIName] = make(map[string]string)
	IndexMap[PPIName]["outerCode"] = "A090105"
	IndexMap[PPIName]["innerCode"] = "HG00021"
	IndexMap[PPIName]["startYear"] = "1978"

	// 工业生产者出厂价格指数(1985=100)
	IndexMap[PPI1Name] = make(map[string]string)
	IndexMap[PPI1Name]["outerCode"] = "A090205"
	IndexMap[PPI1Name]["innerCode"] = "" // TODO: no need yet
	IndexMap[PPI1Name]["startYear"] = "1985"

	// 工业生产者出厂价格指数(上月=100)
	IndexMap[PPI2Name] = make(map[string]string)
	IndexMap[PPI2Name]["outerCode"] = "A01080701"
	IndexMap[PPI2Name]["innerCode"] = "HG00022"
	IndexMap[PPI2Name]["startYear"] = "2011"

	// 工业生产者出厂价格指数(上年同月=100)
	IndexMap[PPI3Name] = make(map[string]string)
	IndexMap[PPI3Name]["outerCode"] = "A01080101"
	IndexMap[PPI3Name]["innerCode"] = "HG00023"
	IndexMap[PPI3Name]["startYear"] = "1996"

	// 工业生产者出厂价格指数(上年同期=100)
	IndexMap[PPI4Name] = make(map[string]string)
	IndexMap[PPI4Name]["outerCode"] = "A01080401"
	IndexMap[PPI4Name]["innerCode"] = "HG00024"
	IndexMap[PPI4Name]["startYear"] = "2011"

	// 居民消费水平
	IndexMap[RCLName] = make(map[string]string)
	IndexMap[RCLName]["outerCode"] = "A020501"
	IndexMap[RCLName]["innerCode"] = "HG00025"
	IndexMap[RCLName]["startYear"] = "1992"

	// 社会消费品零售总额
	IndexMap[SCGName] = make(map[string]string)
	IndexMap[SCGName]["outerCode"] = "A0H01"
	IndexMap[SCGName]["innerCode"] = "HG00026"
	IndexMap[SCGName]["startYear"] = "1952"

	// 社会消费品零售总额_当期值
	IndexMap[SCG1Name] = make(map[string]string)
	IndexMap[SCG1Name]["outerCode"] = "A070101"
	IndexMap[SCG1Name]["innerCode"] = "HG00027"
	IndexMap[SCG1Name]["startYear"] = "1984"

	// 社会消费品零售总额_累计值
	IndexMap[SCG2Name] = make(map[string]string)
	IndexMap[SCG2Name]["outerCode"] = "A070102"
	IndexMap[SCG2Name]["innerCode"] = "HG00028"
	IndexMap[SCG2Name]["startYear"] = "2000"

	// 社会消费品零售总额_同比增长
	IndexMap[SCG3Name] = make(map[string]string)
	IndexMap[SCG3Name]["outerCode"] = "A070103"
	IndexMap[SCG3Name]["innerCode"] = "HG00029"
	IndexMap[SCG3Name]["startYear"] = "2000"

	// 社会消费品零售总额_累计增长
	IndexMap[SCG4Name] = make(map[string]string)
	IndexMap[SCG4Name]["outerCode"] = "A070104"
	IndexMap[SCG4Name]["innerCode"] = "HG00030"
	IndexMap[SCG4Name]["startYear"] = "2000"

	// 城镇居民人均可支配收入
	IndexMap[URIName] = make(map[string]string)
	IndexMap[URIName]["outerCode"] = "A0A0103"
	IndexMap[URIName]["innerCode"] = "HG00031"
	IndexMap[URIName]["startYear"] = "2013"

	// 城镇居民人均可支配收入_同比增长
	IndexMap[URI1Name] = make(map[string]string)
	IndexMap[URI1Name]["outerCode"] = "A0A0104"
	IndexMap[URI1Name]["innerCode"] = "HG00032"
	IndexMap[URI1Name]["startYear"] = "2014"

	// 城镇居民人均可支配收入_累计值
	IndexMap[URI2Name] = make(map[string]string)
	IndexMap[URI2Name]["outerCode"] = "A050201"
	IndexMap[URI2Name]["innerCode"] = "HG00033"
	IndexMap[URI2Name]["startYear"] = "2013"

	// 城镇居民人均可支配收入_累计增长
	IndexMap[URI3Name] = make(map[string]string)
	IndexMap[URI3Name]["outerCode"] = "A050202"
	IndexMap[URI3Name]["innerCode"] = "HG00034"
	IndexMap[URI3Name]["startYear"] = "2014"
}

func init() {
	setUpIndexMap()
}