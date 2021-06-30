package provincecode

// 省份代码

const (
	BeiJingCode      = "110000" // 北京市代码
	TianJinCode      = "120000" // 天津市代码
	HeBeiCode        = "130000" // 河北省代码
	ShanXiCode       = "140000" // 山西省代码
	NeiMengGuCode    = "150000" // 内蒙古自治区代码
	LiaoNingCode     = "210000" // 辽宁省代码
	JiLinCode        = "220000" // 吉林省代码
	HeiLongJiangCode = "230000" // 黑龙江省代码
	ShangHaiCode     = "310000" // 上海市代码
	JiangSuCode      = "320000" // 江苏省代码
	ZheJiangCode     = "330000" // 浙江省代码
	AnHuiCode        = "340000" // 安徽省代码
	FuJianCode       = "350000" // 福建省代码
	JiangXiCode      = "360000" // 江西省代码
	ShanDongCode     = "370000" // 山东省代码
	HeNanCode        = "410000" // 河南省代码
	HuBeiCode        = "420000" // 湖北省代码
	HuNanCode        = "430000" // 湖南省代码
	GuangDongCode    = "440000" // 广东省代码
	GuangXiCode      = "450000" // 广西壮族自治区代码
	HaiNanCode       = "460000" // 海南省代码
	ChongQingCode    = "500000" // 重庆市代码
	SiChuanCode      = "510000" // 四川省代码
	GuiZhouCode      = "520000" // 贵州省代码
	YunNanCode       = "530000" // 云南省代码
	XiZangCode       = "540000" // 西藏自治区代码
	ShanXiiCode      = "610000" // 陕西省代码
	GanSuCode        = "620000" // 甘肃省代码
	QingHaiCode      = "630000" // 青海省代码
	NingXiaCode      = "640000" // 宁夏回族自治区代码
	XinJiangCode     = "650000" // 新疆维吾尔族自治区代码
)

var ProvinceCode map[string]string
var CodeProvince map[string]string

func init() {
	pc := make(map[string]string)
	pc["BeiJingCode"] = BeiJingCode
	pc["TianJinCode"] = TianJinCode
	pc["HeBeiCode"] = HeBeiCode
	pc["ShanXiCode"] = ShanXiCode
	pc["NeiMengGuCode"] = NeiMengGuCode
	pc["LiaoNingCode"] = LiaoNingCode
	pc["JiLinCode"] = JiLinCode
	pc["HeiLongJiangCode"] = HeiLongJiangCode
	pc["ShangHaiCode"] = ShangHaiCode
	pc["JiangSuCode"] = JiangSuCode
	pc["ZheJiangCode"] = ZheJiangCode
	pc["AnHuiCode"] = AnHuiCode
	pc["FuJianCode"] = FuJianCode
	pc["JiangXiCode"] = JiangXiCode
	pc["ShanDongCode"] = ShanDongCode
	pc["HeNanCode"] = HeNanCode
	pc["HuBeiCode"] = HuBeiCode
	pc["HuNanCode"] = HuNanCode
	pc["GuangDongCode"] = GuangDongCode
	pc["GuangXiCode"] = GuangXiCode
	pc["HaiNanCode"] = HaiNanCode
	pc["ChongQingCode"] = ChongQingCode
	pc["SiChuanCode"] = SiChuanCode
	pc["GuiZhouCode"] = GuiZhouCode
	pc["YunNanCode"] = YunNanCode
	pc["XiZangCode"] = XiZangCode
	pc["ShanXiiCode"] = ShanXiiCode
	pc["GanSuCode"] = GanSuCode
	pc["QingHaiCode"] = QingHaiCode
	pc["NingXiaCode"] = NingXiaCode
	pc["XinJiangCode"] = XinJiangCode
	ProvinceCode = pc

	cp := make(map[string]string)
	cp[BeiJingCode] = "北京"
	cp[TianJinCode] = "天津"
	cp[HeBeiCode] = "河北"
	cp[ShanXiCode] = "山西"
	cp[NeiMengGuCode] = "内蒙古"
	cp[LiaoNingCode] = "辽宁"
	cp[JiLinCode] = "吉林"
	cp[HeiLongJiangCode] = "黑龙江"
	cp[ShangHaiCode] = "上海"
	cp[JiangSuCode] = "江苏"
	cp[ZheJiangCode] = "浙江"
	cp[AnHuiCode] = "安徽"
	cp[FuJianCode] = "福建"
	cp[JiangXiCode] = "江西"
	cp[ShanDongCode] = "山东"
	cp[HeNanCode] = "河南"
	cp[HuBeiCode] = "湖北"
	cp[HuNanCode] = "湖南"
	cp[GuangDongCode] = "广东"
	cp[GuangXiCode] = "广西"
	cp[HaiNanCode] = "海南"
	cp[ChongQingCode] = "重庆"
	cp[SiChuanCode] = "四川"
	cp[GuiZhouCode] = "贵州"
	cp[YunNanCode] = "云南"
	cp[XiZangCode] = "西藏"
	cp[ShanXiiCode] = "陕西"
	cp[GanSuCode] = "甘肃"
	cp[QingHaiCode] = "青海"
	cp[NingXiaCode] = "宁夏"
	cp[XinJiangCode] = "新疆"
	CodeProvince = cp
}
