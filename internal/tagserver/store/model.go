package store

type PolicyNewsOrg struct {
	NEWS_GUID              string
	NEWS_ID                int
	NEWS_TS                string
	NEWS_SUMMARY           string
	EMOTION_INDICATOR      string             `json:"新闻舆情指标"`
	EMOTION_INDICATOR_NAME string             `json:"新闻舆情指标名称"`
	EMOTION_WEIGHT         string             `json:"新闻舆情指标权重"`
	EMOTION_DETAIL         string             `json:"三类舆情指标对应权重"`
	RegionMap              map[string]int     `json:"地区信息"`
	CompanyMap             map[string]int     `json:"公司信息"`
	IndustryMap            map[string]float64 `json:"行業信息"`
}

type PolicyNews struct {
	NEWS_GUID              string `json:"新闻ID"`
	EMOTION_INDICATOR      string `json:"新闻舆情指标"`
	EMOTION_INDICATOR_NAME string `json:"新闻舆情指标名称"`
	EMOTION_WEIGHT         string `json:"新闻舆情指标权重"`
	EMOTION_DETAIL         string `json:"三类舆情指标对应权重"`
	NEWS_LABELS            string `json:"新闻标签列表"`
	INDUSTRY_LABELS        string `json:"行业标签列表"`
	COMPANY_LABELS         string `json:"公司标签列表"`
	REGION_LABELS          string `json:"地区标签列表"`
	EVENT_LABELS           string `json:"事件标签列表"`
}

type NewsRegion struct {
	REGION_LABEL_GUID string `json:"新闻地区标签ID"`
	NEWS_GUID         string `json:"新闻ID"`
	REGION_CODE       string `json:"地区代码"`
	REGION_NAME       string `json:"地区名称"`
	ENGLISH_NAME      string `json:"英文名称"`
	NEWS_ID           int    `json:"新闻原始ID"`
	NEWS_TS           string `json:"新闻时间戳"`
	RELEVANCE         string `jaon:"标签相关度"`
	EMOTION_INDICATOR string `jaon:"新闻舆情指标"`
	EMOTION_WEIGHT    string `jaon:"新闻舆情指标权重"`
	EMOTION_DETAIL    string `jaon:"三类舆情指标对应权重"`
}

type NewsCompany struct {
	COMP_LABEL_GUID   string `json:"新闻公司标签ID"`
	NEWS_GUID         string `json:"新闻ID"`
	STOCK_CODE        string `json:"股票代码"`
	CREDIT_CODE       string `json:"统一社会信息代码"`
	COMP_GUID         string `json:"系统公司ID"`
	COMPANY_ID        string `json:"数据源公司ID"`
	CHINESE_NAME      string `json:"中文名称"`
	ENGLISH_NAME      string `json:"英文名称"`
	NEWS_ID           int    `json:"新闻原始ID"`
	NEWS_TS           string `json:"新闻时间戳"`
	RELEVANCE         string `json:"标签相关度"`
	EMOTION_INDICATOR string `json:"新闻舆情指标"`
	EMOTION_WEIGHT    string `json:"新闻舆情指标权重"`
	EMOTION_DETAIL    string `json:"三类舆情指标对应权重"`
}

type NewsIndustry struct {
	INDUSTRY_LABEL_GUID string `json:"新闻行业标签ID"`
	NEWS_GUID           string `json:"新闻ID"`
	INDUSTRY_CODE       string `json:"行业代码"`
	INDUSTRY_NAME       string `json:"行业名称"`
	ENGLISH_NAME        string `json:"英文名称"`
	NEWS_ID             int    `json:"新闻原始ID"`
	NEWS_TS             string `json:"新闻时间戳"`
	RELEVANCE           string `json:"标签相关度"`
	EMOTION_INDICATOR   string `json:"新闻舆情指标"`
	EMOTION_WEIGHT      string `json:"新闻舆情指标权重"`
	EMOTION_DETAIL      string `json:"三类舆情指标对应权重"`
}
