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

// 指标代码

const (
	CPICode  = "A090101"   // 居民消费价格指数(上年=100)[start: 1951]<run>
	CPI1Code = "A090201"   // 居民消费价格指数(1978=100)[start: 1978]
	CPI2Code = "A01030101" // 居民消费价格指数(上月=100)[start: 2016]
	CPI3Code = "A01010101" // 居民消费价格指数(上年同月=100)[start: 2016]<run>
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
