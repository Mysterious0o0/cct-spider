package findregion

import (
	"regexp"
)

func FindRegion(regionPat, s string) (region string) {
	reg := regexp.MustCompile(regionPat)
	regions := reg.FindStringSubmatch(s)
	if len(regions) == 0 {
		return
	}
	return regions[0]
}

func FindAllRegion(regionPat, s string) (region []string) {
	reg := regexp.MustCompile(regionPat)
	regions := reg.FindAllStringSubmatch(s, -1)
	if len(regions) == 0 {
		return
	}
	for _, r := range regions{
		region = append(region, r[0])
	}
	return
}