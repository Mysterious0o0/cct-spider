package indexcode

// 指标起始年

const (
	CPIStartYear  = 1951
	CPI1StartYear = 1978
	CPI2StartYear = 2016
	CPI3StartYear = 2016
	CPI4StartYear = 2016
	CQCStartYear  = 1991
	CQC1StartYear = 1999
	CQC2StartYear = 1999
	FAIStartYear  = 1990
	FAI1StartYear = 1990
	FAI2StartYear = 1998
	FAI3StartYear = 2007
	GDPStartYear  = 1952
	GDPRStartYear = 1992
	GDP1StartYear = 1992
	GDP2StartYear = 1992
	HCEStartYear  = 2013
	HCE1StartYear = 2014
	HCE2StartYear = 2013
	HCE3StartYear = 2014
	IAVStartYear  = 1952
	IAV1StartYear = 1998
	IAV2StartYear = 1998
	IAV3StartYear = 1992
	IAV4StartYear = 1992
	PMIStartYear  = 2005
	PPIStartYear  = 1978
	PPI1StartYear = 1985
	PPI2StartYear = 2011
	PPI3StartYear = 1996
	PPI4StartYear = 2011
	RCLStartYear  = 1992
	SCGStartYear  = 1952
	SCG1StartYear = 1984
	SCG2StartYear = 2000
	SCG3StartYear = 2000
	SCG4StartYear = 2000
	URIStartYear  = 2013
	URI1StartYear = 2014
	URI2StartYear = 2013
	URI3StartYear = 2014
)

// 国家统计局指标代码

const (
	CPICode  = "A090101"   // 居民消费价格指数(上年=100)[start: 1951]
	CPI1Code = "A090201"   // 居民消费价格指数(1978=100)[start: 1978]
	CPI2Code = "A01030101" // 居民消费价格指数(上月=100)[start: 2016]
	CPI3Code = "A01010101" // 居民消费价格指数(上年同月=100)[start: 2016]
	CPI4Code = "A01020101" // 居民消费价格指数(上年同期=100)[start: 2016]
	CQCCode  = "A0L0309"   // 货币和准货币(M2)供应量同比增长率(%)[start: 1991]
	CQC1Code = "A0D0101"   // 货币和准货币(M2)供应量_期末值(亿元)[start: 1999]
	CQC2Code = "A0D0102"   // 货币和准货币(M2)供应量_同比增长(%)[start: 1999]
	FAICode  = "A090107"   // 固定资产投资价格指数(上年=100)[start: 1990]
	FAI1Code = "A090207"   // 固定资产投资价格指数(1990=100)[start: 1990]
	FAI2Code = "A060201"   // 固定资产投资价格指数_当季值(上年同季=100)[start: 1998]
	FAI3Code = "A060301"   // 固定资产投资价格指数_累计值(上年同期=100)[start: 2007]
	GDPCode  = "A020102"   // 国内生产总值(亿元)[start: 1952]
	GDPRCode = "A020101"   // 地区生产总值(亿元)[start: 1992]
	GDP1Code = "A010101"   // 国内生产总值_当季值(亿元)[start: 1992]
	GDP2Code = "A010102"   // 国内生产总值_累计值(亿元)[start: 1992]
	HCECode  = "A0A0107"   // 居民人均消费支出(元)[start: 2013]
	HCE1Code = "A0A0108"   // 居民人均消费支出_同比增长(%)[start: 2014]
	HCE2Code = "A050109"   // 居民人均消费支出_累计值(元)[start: 2013]
	HCE3Code = "A05010A"   // 居民人均消费支出_累计增长(%)[start: 2014]
	IAVCode  = "A020403"   // 工业增加值(亿元)[start: 1952]
	IAV1Code = "A020101"   // 工业增加值_同比增长(%)[start: 1998]
	IAV2Code = "A020102"   // 工业增加值_累计增长(%)[start: 1998]
	IAV3Code = "A01010B"   // 工业增加值_当季值(亿元)[start: 1992]
	IAV4Code = "A01010C"   // 工业增加值_累计值(亿元)[start: 1992]
	PMICode  = "A0B0101"   // 制造业采购经理指数(%)[start: 2005]
	PPICode  = "A090105"   // 工业生产者出厂价格指数(上年=100)[start: 1978]
	PPI1Code = "A090205"   // 工业生产者出厂价格指数(1985=100)[start: 1985]
	PPI2Code = "A01080701" // 工业生产者出厂价格指数(上月=100)[start: 2011]
	PPI3Code = "A01080101" // 工业生产者出厂价格指数(上年同月=100)[start: 1996]
	PPI4Code = "A01080401" // 工业生产者出厂价格指数(上年同期=100)[start: 2011]
	RCLCode  = "A020501"   // 居民消费水平(元)[start: 1992]
	SCGCode  = "A0H01"     // 社会消费品零售总额(亿元)[start: 1952]
	SCG1Code = "A070101"   // 社会消费品零售总额_当期值(亿元)[start: 1984]
	SCG2Code = "A070102"   // 社会消费品零售总额_累计值(亿元)[start: 2000]
	SCG3Code = "A070103"   // 社会消费品零售总额_同比增长(%)[start: 2000]
	SCG4Code = "A070104"   // 社会消费品零售总额_累计增长(%)[start: 2000]
	URICode  = "A0A0103"   // 城镇居民人均可支配收入(元)[start: 2013]
	URI1Code = "A0A0104"   // 城镇居民人均可支配收入_同比增长(%)[start: 2014]
	URI2Code = "A050201"   // 城镇居民人均可支配收入_累计值(元)[start: 2013]
	URI3Code = "A050202"   // 城镇居民人均可支配收入_累计增长(%)[start: 2014]
)

// 指标名称(国家统计局和内部统一)

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

// 系统内部指标代码

const (
	InnerCPICode  = "HG00003" // 居民消费价格指数(上年=100)[start: 1951]
	InnerCPI1Code = ""        // 居民消费价格指数(1978=100)[start: 1978]
	InnerCPI2Code = ""        // 居民消费价格指数(上月=100)[start: 2016]
	InnerCPI3Code = "HG00004" // 居民消费价格指数(上年同月=100)[start: 2016]
	InnerCPI4Code = ""        // 居民消费价格指数(上年同期=100)[start: 2016]
	InnerCQCCode  = "HG00005" // 货币和准货币(M2)供应量同比增长率(%)[start: 1991]
	InnerCQC1Code = "HG00006" // 货币和准货币(M2)供应量_期末值(亿元)[start: 1999]
	InnerCQC2Code = "HG00007" // 货币和准货币(M2)供应量_同比增长(%)[start: 1999]
	InnerFAICode  = "HG00008" // 固定资产投资价格指数(上年=100)[start: 1990]
	InnerFAI1Code = ""        // 固定资产投资价格指数(1990=100)[start: 1990]
	InnerFAI2Code = "HG00009" // 固定资产投资价格指数_当季值(上年同季=100)[start: 1998]
	InnerFAI3Code = "HG00010" // 固定资产投资价格指数_累计值(上年同期=100)[start: 2007]
	InnerGDPCode  = "HG00001" // 国内生产总值(亿元)[start: 1952]
	InnerGDPRCode = "HG00002" // 地区生产总值(亿元)[start: 1992]
	InnerGDP1Code = ""        // 国内生产总值_当季值(亿元)[start: 1992]
	InnerGDP2Code = ""        // 国内生产总值_累计值(亿元)[start: 1992]
	InnerHCECode  = "HG00011" // 居民人均消费支出(元)[start: 2013]
	InnerHCE1Code = "HG00012" // 居民人均消费支出_同比增长(%)[start: 2014]
	InnerHCE2Code = "HG00013" // 居民人均消费支出_累计值(元)[start: 2013]
	InnerHCE3Code = "HG00014" // 居民人均消费支出_累计增长(%)[start: 2014]
	InnerIAVCode  = "HG00015" // 工业增加值(亿元)[start: 1952]
	InnerIAV1Code = "HG00016" // 工业增加值_同比增长(%)[start: 1998]
	InnerIAV2Code = "HG00017" // 工业增加值_累计增长(%)[start: 1998]
	InnerIAV3Code = "HG00018" // 工业增加值_当季值(亿元)[start: 1992]
	InnerIAV4Code = "HG00019" // 工业增加值_累计值(亿元)[start: 1992]
	InnerPMICode  = "HG00020" // 制造业采购经理指数(%)[start: 2005]
	InnerPPICode  = "HG00021" // 工业生产者出厂价格指数(上年=100)[start: 1978]
	InnerPPI1Code = ""        // 工业生产者出厂价格指数(1985=100)[start: 1985]
	InnerPPI2Code = "HG00022" // 工业生产者出厂价格指数(上月=100)[start: 2011]
	InnerPPI3Code = "HG00023" // 工业生产者出厂价格指数(上年同月=100)[start: 1996]
	InnerPPI4Code = "HG00024" // 工业生产者出厂价格指数(上年同期=100)[start: 2011]
	InnerRCLCode  = "HG00025" // 居民消费水平(元)[start: 1992]
	InnerSCGCode  = "HG00026" // 社会消费品零售总额(亿元)[start: 1952]
	InnerSCG1Code = "HG00027" // 社会消费品零售总额_当期值(亿元)[start: 1984]
	InnerSCG2Code = "HG00028" // 社会消费品零售总额_累计值(亿元)[start: 2000]
	InnerSCG3Code = "HG00029" // 社会消费品零售总额_同比增长(%)[start: 2000]
	InnerSCG4Code = "HG00030" // 社会消费品零售总额_累计增长(%)[start: 2000]
	InnerURICode  = "HG00031" // 城镇居民人均可支配收入(元)[start: 2013]
	InnerURI1Code = "HG00032" // 城镇居民人均可支配收入_同比增长(%)[start: 2014]
	InnerURI2Code = "HG00033" // 城镇居民人均可支配收入_累计值(元)[start: 2013]
	InnerURI3Code = "HG00034" // 城镇居民人均可支配收入_累计增长(%)[start: 2014]
)

var StatInner map[string]string // 国家统计局指标代码：内部指标代码
var CodeName map[string]string  // 国家统计局指标代码：指标名称

func init() {
	StatInner = make(map[string]string)
	StatInner[CPICode] = InnerCPICode
	StatInner[CPI1Code] = InnerCPI1Code
	StatInner[CPI2Code] = InnerCPI2Code
	StatInner[CPI3Code] = InnerCPI3Code
	StatInner[CPI4Code] = InnerCPI4Code
	StatInner[CQCCode] = InnerCQCCode
	StatInner[CQC1Code] = InnerCQC1Code
	StatInner[CQC2Code] = InnerCQC2Code
	StatInner[FAICode] = InnerFAICode
	StatInner[FAI1Code] = InnerFAI1Code
	StatInner[FAI2Code] = InnerFAI2Code
	StatInner[FAI3Code] = InnerFAI3Code
	StatInner[GDPCode] = InnerGDPCode
	StatInner[GDPRCode] = InnerGDPRCode
	StatInner[GDP1Code] = InnerGDP1Code
	StatInner[GDP2Code] = InnerGDP2Code
	StatInner[HCECode] = InnerHCECode
	StatInner[HCE1Code] = InnerHCE1Code
	StatInner[HCE2Code] = InnerHCE2Code
	StatInner[HCE3Code] = InnerHCE3Code
	StatInner[IAVCode] = InnerIAVCode
	StatInner[IAV1Code] = InnerIAV1Code
	StatInner[IAV2Code] = InnerIAV2Code
	StatInner[IAV3Code] = InnerIAV3Code
	StatInner[IAV4Code] = InnerIAV4Code
	StatInner[PMICode] = InnerPMICode
	StatInner[PPICode] = InnerPPICode
	StatInner[PPI1Code] = InnerPPI1Code
	StatInner[PPI2Code] = InnerPPI2Code
	StatInner[PPI3Code] = InnerPPI3Code
	StatInner[PPI4Code] = InnerPPI4Code
	StatInner[RCLCode] = InnerRCLCode
	StatInner[SCGCode] = InnerSCGCode
	StatInner[SCG1Code] = InnerSCG1Code
	StatInner[SCG2Code] = InnerSCG2Code
	StatInner[SCG3Code] = InnerSCG3Code
	StatInner[SCG4Code] = InnerSCG4Code
	StatInner[URICode] = InnerURICode
	StatInner[URI1Code] = InnerURI1Code
	StatInner[URI2Code] = InnerURI2Code
	StatInner[URI3Code] = InnerURI3Code

	CodeName = make(map[string]string)
	CodeName[CPICode] = CPIName
	CodeName[CPI1Code] = CPI1Name
	CodeName[CPI2Code] = CPI2Name
	CodeName[CPI3Code] = CPI3Name
	CodeName[CPI4Code] = CPI4Name
	CodeName[CQCCode] = CQCName
	CodeName[CQC1Code] = CQC1Name
	CodeName[CQC2Code] = CQC2Name
	CodeName[FAICode] = FAIName
	CodeName[FAI1Code] = FAI1Name
	CodeName[FAI2Code] = FAI2Name
	CodeName[FAI3Code] = FAI3Name
	CodeName[GDPCode] = GDPName
	CodeName[GDPRCode] = GDPRName
	CodeName[GDP1Code] = GDP1Name
	CodeName[GDP2Code] = GDP2Name
	CodeName[HCECode] = HCEName
	CodeName[HCE1Code] = HCE1Name
	CodeName[HCE2Code] = HCE2Name
	CodeName[HCE3Code] = HCE3Name
	CodeName[IAVCode] = IAVName
	CodeName[IAV1Code] = IAV1Name
	CodeName[IAV2Code] = IAV2Name
	CodeName[IAV3Code] = IAV3Name
	CodeName[IAV4Code] = IAV4Name
	CodeName[PMICode] = PMIName
	CodeName[PPICode] = PPIName
	CodeName[PPI1Code] = PPI1Name
	CodeName[PPI2Code] = PPI2Name
	CodeName[PPI3Code] = PPI3Name
	CodeName[PPI4Code] = PPI4Name
	CodeName[RCLCode] = RCLName
	CodeName[SCGCode] = SCGName
	CodeName[SCG1Code] = SCG1Name
	CodeName[SCG2Code] = SCG2Name
	CodeName[SCG3Code] = SCG3Name
	CodeName[SCG4Code] = SCG4Name
	CodeName[URICode] = URIName
	CodeName[URI1Code] = URI1Name
	CodeName[URI2Code] = URI2Name
	CodeName[URI3Code] = URI3Name
}
