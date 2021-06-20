package mca

import (
	"testing"
)

type a struct {
	Test string `json:"wh"`
}

func TestXpath(t *testing.T) {
	infoUrl := "http://xxgk.mca.gov.cn:8011/gdnps/searchIndex.jsp?params=%257B%2522goPage%2522%253A1%252C%2522orderBy%2522%253A%255B%257B%2522orderBy%2522%253A%2522scrq%2522%252C%2522reverse%2522%253Atrue%257D%252C%257B%2522orderBy%2522%253A%2522orderTime%2522%252C%2522reverse%2522%253Atrue%257D%255D%252C%2522pageSize%2522%253A5000%252C%2522queryParam%2522%253A%255B%257B%2522shortName%2522%253A%2522ownSubjectDn%2522%252C%2522value%2522%253A%2522%252F1%252F3%252F164%252F230%2522%257D%252C%257B%2522shortName%2522%253A%2522fbjg%2522%252C%2522value%2522%253A%2522%252F1%252F3%252F164%252F230%2522%257D%252C%257B%257D%252C%257B%257D%255D%252C%2522doRepeated%2522%253A0%257D"
	GetHtmlInfo(infoUrl)
}
