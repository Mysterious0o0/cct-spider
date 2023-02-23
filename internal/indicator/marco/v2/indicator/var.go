package indicator

var (
	table = "T_DMAA_BASE_TARGET_VALUE"
)

var (
	MarcoYearSQL = "SELECT DISTINCT ACCT_YEAR FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR)) AND " +
		"TARGET_CODE = '%s'"
	MarcoSeasonSQL = "SELECT DISTINCT ACCT_YEAR FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR, ACCT_QUARTOR)) AND " +
		"TARGET_CODE = '%s'"
	MarcoMonthSQL = "SELECT DISTINCT ACCT_YEAR FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR, ACCT_MONTH)) AND " +
		"TARGET_CODE = '%s'"
	ProvinceYearSQL = "SELECT DISTINCT ACCT_YEAR FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR, REGION_CODE)) AND " +
		"TARGET_CODE = '%s' AND " +
		"REGION_CODE != ''"
	ProvinceSeasonSQL = "SELECT DISTINCT ACCT_YEAR FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR, ACCT_QUARTOR, REGION_CODE)) AND " +
		"TARGET_CODE = '%s' AND " +
		"REGION_CODE != ''"
	ProvinceMonthSQL = "SELECT DISTINCT ACCT_YEAR FROM %s WHERE " +
		"VALUE_GUID = MD5(CONCAT(TARGET_CODE, ACCT_YEAR, ACCT_MONTH, REGION_CODE)) AND " +
		"TARGET_CODE = '%s' AND " +
		"REGION_CODE != ''"
)