package main

import (
	"github.com/xiaogogonuo/cct-spider/internal/indicator"
	cpiNationMonth "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/cpi/nation/month"
	cpiNationYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/cpi/nation/year"
	cpiRegionMonth "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/cpi/region/month"
	cpiRegionYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/cpi/region/year"
	cqcNationMonth "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/cqc/nation/month"
	cqcNationYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/cqc/nation/year"
	faiNationSeason "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/fai/nation/season"
	faiNationYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/fai/nation/year"

	gdpNationSeason "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/gdp/nation/season"
	gdpNationYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/gdp/nation/year"
	gdpRegionYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/gdp/region/year"

	hecNationSeason "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/hce/nation/season"
	hecNationYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/hce/nation/year"

	iavNationMonth "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/iav/nation/month"
	iavNationSeason "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/iav/nation/season"
	iavNationYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/iav/nation/year"

	pmiNationMonth "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/pmi/nation/month"

	ppiNationMonth "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/ppi/nation/month"
	ppiNationYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/ppi/nation/year"

	rclRegionYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/rcl/region/year"

	scgNationMonth "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/scg/nation/month"
	scgNationYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/scg/nation/year"

	uriNationSeason "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/uri/nation/season"
	uriNationYear "github.com/xiaogogonuo/cct-spider/internal/stat/api/v1/uri/nation/year"
)

const (
	cpiNationMonthRun  = false
	cpiNationYearRun   = false
	cpiRegionMonthRun  = false
	cpiRegionYearRun   = false

	cqcNationMonthRun  = false
	cqcNationYearRun   = false

	faiNationSeasonRun = false
	faiNationYearRun   = false

	gdpNationSeasonRun = false
	gdpNationYearRun   = false
	gdpRegionYearRun   = false

	hecNationSeasonRun = false
	hecNationYearRun   = false

	iavNationMonthRun  = false
	iavNationSeasonRun = false
	iavNationYearRun   = false

	pmiNationMonthRun  = false

	ppiNationMonthRun  = false
	ppiNationYearRun   = false

	rclRegionYearRun   = false

	scgNationMonthRun  = false
	scgNationYearRun   = false

	uriNationSeasonRun = false
	uriNationYearRun   = false
)

func main() {
	// CPI
	{
		if cpiNationMonthRun {
			cpiNationMonth.Run()
		}
		if cpiNationYearRun {
			cpiNationYear.Run()
		}
		if cpiRegionMonthRun {
			cpiRegionMonth.Run()
		}
		if cpiRegionYearRun {
			cpiRegionYear.Run()
		}
	}

	// CQC
	{
		if cqcNationMonthRun {
			cqcNationMonth.Run()
		}
		if cqcNationYearRun {
			cqcNationYear.Run()
		}
	}

	// FAI
	{
		if faiNationSeasonRun {
			faiNationSeason.Run()
		}
		if faiNationYearRun {
			faiNationYear.Run()
		}
	}

	// GDP
	{
		if gdpNationSeasonRun {
			gdpNationSeason.Run()
		}
		if gdpNationYearRun {
			gdpNationYear.Run()
		}
		if gdpRegionYearRun {
			gdpRegionYear.Run()
		}
	}

	// HCE
	{
	    if hecNationSeasonRun {
			hecNationSeason.Run()
		}
		if hecNationYearRun {
			hecNationYear.Run()
		}
	}

	// IAV
	{
		if iavNationMonthRun {
			iavNationMonth.Run()
		}
		if iavNationSeasonRun {
			iavNationSeason.Run()
		}
		if iavNationYearRun {
			iavNationYear.Run()
		}
	}

	// PMI
	{
		if pmiNationMonthRun {
			pmiNationMonth.Run()
		}
	}

	// PPI
	{
		if ppiNationMonthRun {
			ppiNationMonth.Run()
		}
		if ppiNationYearRun {
			ppiNationYear.Run()
		}
	}

	// RCL
	{
		if rclRegionYearRun {
			rclRegionYear.Run()
		}
	}

	// SCG
	{
		if scgNationMonthRun {
			scgNationMonth.Run()
		}
		if scgNationYearRun {
			scgNationYear.Run()
		}
	}

	// URI
	{
		if uriNationSeasonRun {
			uriNationSeason.Run()
		}
		if uriNationYearRun {
			uriNationYear.Run()
		}
	}

	indicator.Start()
}