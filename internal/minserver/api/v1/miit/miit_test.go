package miit

import (
	"testing"
)

//func TestBody(t *testing.T) {
//	s := `{
//   "roles": null,
//   "permissions": null,
//   "message": "查询成功",
//   "data": {
//       "searchResult": {
//           "dataResults": [
//               {
//                   "pkValue": null,
//                   "score": 0,
//                   "docId": 0,
//                   "groupCount": 1,
//                   "groupValue": "73831152",
//                   "similarRate": 0,
//                   "groupData": [
//                       {
//                           "pkValue": "34a594d2b9bf4fba9141baf1d929a15e",
//                           "score": "NaN",
//                           "docId": 0,
//                           "groupCount": 0,
//                           "groupValue": "",
//                           "similarRate": 0,
//                           "groupData": [],
//                           "data": {
//                               "title": "工业和信息化部办公厅关于开展车联网身份认证和安全信任试点工作的通知",
//                               "url": "/zwgk/zcwj/wjfb/qt/art/2021/art_34a594d2b9bf4fba9141baf1d929a15e.html",
//                               "jsearch_url": "visit/link?url=%2Fzwgk%2Fzcwj%2Fwjfb%2Fqt%2Fart%2F2021%2Fart_34a594d2b9bf4fba9141baf1d929a15e.html&websiteid=110000000000000&cateid=57&infoid=34a594d2b9bf4fba9141baf1d929a15e"
//                           },
//                           "tableName": "cms_info"
//                       }
//                   ],
//                   "data": {
//                       "_value": "73831152",
//                       "_count": "1",
//                       "max_deploytime": "1623306779744"
//                   },
//                   "tableName": null
//               },
//               {
//                   "pkValue": null,
//                   "score": 0,
//                   "docId": 0,
//                   "groupCount": 1,
//                   "groupValue": "-1215871732",
//                   "similarRate": 0,
//                   "groupData": [
//                       {
//                           "pkValue": "41ae7b048792416ca167bde07368475f",
//                           "score": "NaN",
//                           "docId": 0,
//                           "groupCount": 0,
//                           "groupValue": "",
//                           "similarRate": 0,
//                           "groupData": [],
//                           "data": {
//                               "title": "四部委关于印发汽车产品生产者责任延伸试点实施方案的通知",
//                               "url": "/zwgk/zcwj/wjfb/zbgy/art/2021/art_41ae7b048792416ca167bde07368475f.html",
//                               "jsearch_url": "visit/link?url=%2Fzwgk%2Fzcwj%2Fwjfb%2Fzbgy%2Fart%2F2021%2Fart_41ae7b048792416ca167bde07368475f.html&websiteid=110000000000000&cateid=57&infoid=41ae7b048792416ca167bde07368475f"
//                           },
//                           "tableName": "cms_info"
//                       }
//                   ],
//                   "data": {
//                       "_value": "-1215871732",
//                       "_count": "1",
//                       "max_deploytime": "1623224303365"
//                   },
//                   "tableName": null
//               },
//               {
//                   "pkValue": null,
//                   "score": 0,
//                   "docId": 0,
//                   "groupCount": 1,
//                   "groupValue": "-424430758",
//                   "similarRate": 0,
//                   "groupData": [
//                       {
//                           "pkValue": "3d20813aa09d4eee82f83ea1def68fa6",
//                           "score": "NaN",
//                           "docId": 0,
//                           "groupCount": 0,
//                           "groupValue": "",
//                           "similarRate": 0,
//                           "groupData": [],
//                           "data": {
//                               "title": "《电信设施保护知识问答》图解",
//                               "url": "/zwgk/zcjd/art/2021/art_3d20813aa09d4eee82f83ea1def68fa6.html",
//                               "jsearch_url": "visit/link?url=%2Fzwgk%2Fzcjd%2Fart%2F2021%2Fart_3d20813aa09d4eee82f83ea1def68fa6.html&websiteid=110000000000000&cateid=57&infoid=3d20813aa09d4eee82f83ea1def68fa6"
//                           },
//                           "tableName": "cms_info"
//                       }
//                   ],
//                   "data": {
//                       "_value": "-424430758",
//                       "_count": "1",
//                       "max_deploytime": "1623113969525"
//                   },
//                   "tableName": null
//               },
//               {
//                   "pkValue": null,
//                   "score": 0,
//                   "docId": 0,
//                   "groupCount": 1,
//                   "groupValue": "-1222510281",
//                   "similarRate": 0,
//                   "groupData": [
//                       {
//                           "pkValue": "30c7489e3b3447318188d074f9f81044",
//                           "score": "NaN",
//                           "docId": 0,
//                           "groupCount": 0,
//                           "groupValue": "",
//                           "similarRate": 0,
//                           "groupData": [],
//                           "data": {
//                               "title": "《关于加快推动区块链技术应用和产业发展的指导意见》解读",
//                               "url": "/zwgk/zcjd/art/2021/art_30c7489e3b3447318188d074f9f81044.html",
//                               "jsearch_url": "visit/link?url=%2Fzwgk%2Fzcjd%2Fart%2F2021%2Fart_30c7489e3b3447318188d074f9f81044.html&websiteid=110000000000000&cateid=57&infoid=30c7489e3b3447318188d074f9f81044"
//                           },
//                           "tableName": "cms_info"
//                       }
//                   ],
//                   "data": {
//                       "_value": "-1222510281",
//                       "_count": "1",
//                       "max_deploytime": "1623112411216"
//                   },
//                   "tableName": null
//               },
//               {
//                   "pkValue": null,
//                   "score": 0,
//                   "docId": 0,
//                   "groupCount": 1,
//                   "groupValue": "1226260589",
//                   "similarRate": 0,
//                   "groupData": [
//                       {
//                           "pkValue": "4646bffe65a348b59e7dd3e52f163ed4",
//                           "score": "NaN",
//                           "docId": 0,
//                           "groupCount": 0,
//                           "groupValue": "",
//                           "similarRate": 0,
//                           "groupData": [],
//                           "data": {
//                               "title": "一图读懂——工业互联网专项工作组2021年工作计划",
//                               "url": "/zwgk/zcjd/art/2021/art_4646bffe65a348b59e7dd3e52f163ed4.html",
//                               "jsearch_url": "visit/link?url=%2Fzwgk%2Fzcjd%2Fart%2F2021%2Fart_4646bffe65a348b59e7dd3e52f163ed4.html&websiteid=110000000000000&cateid=57&infoid=4646bffe65a348b59e7dd3e52f163ed4"
//                           },
//                           "tableName": "cms_info"
//                       }
//                   ],
//                   "data": {
//                       "_value": "1226260589",
//                       "_count": "1",
//                       "max_deploytime": "1623049593394"
//                   },
//                   "tableName": null
//               },
//               {
//                   "pkValue": null,
//                   "score": 0,
//                   "docId": 0,
//                   "groupCount": 1,
//                   "groupValue": "-1262862104",
//                   "similarRate": 0,
//                   "groupData": [
//                       {
//                           "pkValue": "a02effb156344a408e8ca5d60d0442de",
//                           "score": "NaN",
//                           "docId": 0,
//                           "groupCount": 0,
//                           "groupValue": "",
//                           "similarRate": 0,
//                           "groupData": [],
//                           "data": {
//                               "title": "关于印发《工业互联网专项工作组2021年工作计划》的通知",
//                               "url": "/zwgk/zcwj/wjfb/txy/art/2021/art_a02effb156344a408e8ca5d60d0442de.html",
//                               "jsearch_url": "visit/link?url=%2Fzwgk%2Fzcwj%2Fwjfb%2Ftxy%2Fart%2F2021%2Fart_a02effb156344a408e8ca5d60d0442de.html&websiteid=110000000000000&cateid=57&infoid=a02effb156344a408e8ca5d60d0442de"
//                           },
//                           "tableName": "cms_info"
//                       }
//                   ],
//                   "data": {
//                       "_value": "-1262862104",
//                       "_count": "1",
//                       "max_deploytime": "1623048669665"
//                   },
//                   "tableName": null
//               },
//               {
//                   "pkValue": null,
//                   "score": 0,
//                   "docId": 0,
//                   "groupCount": 1,
//                   "groupValue": "1085401905",
//                   "similarRate": 0,
//                   "groupData": [
//                       {
//                           "pkValue": "851f2059f13d41a8bba59c8dce9401a8",
//                           "score": "NaN",
//                           "docId": 0,
//                           "groupCount": 0,
//                           "groupValue": "",
//                           "similarRate": 0,
//                           "groupData": [],
//                           "data": {
//                               "title": "两部门关于加快推动区块链技术应用和产业发展的指导意见",
//                               "url": "/zwgk/zcwj/wjfb/rjy/art/2021/art_851f2059f13d41a8bba59c8dce9401a8.html",
//                               "jsearch_url": "visit/link?url=%2Fzwgk%2Fzcwj%2Fwjfb%2Frjy%2Fart%2F2021%2Fart_851f2059f13d41a8bba59c8dce9401a8.html&websiteid=110000000000000&cateid=57&infoid=851f2059f13d41a8bba59c8dce9401a8"
//                           },
//                           "tableName": "cms_info"
//                       }
//                   ],
//                   "data": {
//                       "_value": "1085401905",
//                       "_count": "1",
//                       "max_deploytime": "1623037684711"
//                   },
//                   "tableName": null
//               },
//               {
//                   "pkValue": null,
//                   "score": 0,
//                   "docId": 0,
//                   "groupCount": 1,
//                   "groupValue": "-1723263363",
//                   "similarRate": 0,
//                   "groupData": [
//                       {
//                           "pkValue": "cb5d3173749f43bc8336c089790244c1",
//                           "score": "NaN",
//                           "docId": 0,
//                           "groupCount": 0,
//                           "groupValue": "",
//                           "similarRate": 0,
//                           "groupData": [],
//                           "data": {
//                               "title": "推进工业文化发展实施方案（2021-2025年）解读",
//                               "url": "/zwgk/zcjd/art/2021/art_cb5d3173749f43bc8336c089790244c1.html",
//                               "jsearch_url": "visit/link?url=%2Fzwgk%2Fzcjd%2Fart%2F2021%2Fart_cb5d3173749f43bc8336c089790244c1.html&websiteid=110000000000000&cateid=57&infoid=cb5d3173749f43bc8336c089790244c1"
//                           },
//                           "tableName": "cms_info"
//                       }
//                   ],
//                   "data": {
//                       "_value": "-1723263363",
//                       "_count": "1",
//                       "max_deploytime": "1622792727336"
//                   },
//                   "tableName": null
//               },
//               {
//                   "pkValue": null,
//                   "score": 0,
//                   "docId": 0,
//                   "groupCount": 1,
//                   "groupValue": "384641526",
//                   "similarRate": 0,
//                   "groupData": [
//                       {
//                           "pkValue": "a25ffe2754bb488da049f257263b6518",
//                           "score": "NaN",
//                           "docId": 0,
//                           "groupCount": 0,
//                           "groupValue": "",
//                           "similarRate": 0,
//                           "groupData": [],
//                           "data": {
//                               "title": "八部门关于印发《推进工业文化发展实施方案（2021-2025年）》的通知",
//                               "url": "/zwgk/zcwj/wjfb/zh/art/2021/art_a25ffe2754bb488da049f257263b6518.html",
//                               "jsearch_url": "visit/link?url=%2Fzwgk%2Fzcwj%2Fwjfb%2Fzh%2Fart%2F2021%2Fart_a25ffe2754bb488da049f257263b6518.html&websiteid=110000000000000&cateid=57&infoid=a25ffe2754bb488da049f257263b6518"
//                           },
//                           "tableName": "cms_info"
//                       }
//                   ],
//                   "data": {
//                       "_value": "384641526",
//                       "_count": "1",
//                       "max_deploytime": "1622786388236"
//                   },
//                   "tableName": null
//               },
//               {
//                   "pkValue": null,
//                   "score": 0,
//                   "docId": 0,
//                   "groupCount": 1,
//                   "groupValue": "1757452896",
//                   "similarRate": 0,
//                   "groupData": [
//                       {
//                           "pkValue": "5b4564af620b4e3d834fbf49570dffd7",
//                           "score": "NaN",
//                           "docId": 0,
//                           "groupCount": 0,
//                           "groupValue": "",
//                           "similarRate": 0,
//                           "groupData": [],
//                           "data": {
//                               "title": "工业和信息化部关于公布2021年换发电子认证服务许可证企业名单（第二批）的通告",
//                               "url": "/zwgk/zcwj/wjfb/tg/art/2021/art_5b4564af620b4e3d834fbf49570dffd7.html",
//                               "jsearch_url": "visit/link?url=%2Fzwgk%2Fzcwj%2Fwjfb%2Ftg%2Fart%2F2021%2Fart_5b4564af620b4e3d834fbf49570dffd7.html&websiteid=110000000000000&cateid=57&infoid=5b4564af620b4e3d834fbf49570dffd7"
//                           },
//                           "tableName": "cms_info"
//                       }
//                   ],
//                   "data": {
//                       "_value": "1757452896",
//                       "_count": "1",
//                       "max_deploytime": "1622776798392"
//                   },
//                   "tableName": null
//               }
//           ],
//           "total": 4855,
//           "totalHits": 4951,
//           "costTime": 59,
//           "times": 0
//       }
//   },
//   "success": true,
//   "code": "200"
//}`
//   var b JsonMiit
//	err := json.Unmarshal([]byte(s), &b)
//	if err != nil {
//		fmt.Println(err)
//	}
//	//fmt.Println(b.Data)
//	for i, v := range b.Data.SearchResult.DataResults{
//		fmt.Println(i, v.GroupData[0].Details.Url)
//	}
//	//fmt.Println(append(b.Data.SearchResult.DataResults))
//}

func TestXpath(t *testing.T) {
	url := "https://www.miit.gov.cn/search-front-server/api/search/info?websiteid=110000000000000&scope=basic&q=&pg=10&cateid=57&pos=title_text%2Cinfocontent%2Ctitlepy&selectFields=title,url,&group=distinct&level=6&sortFields=%5B%7B%22name%22%3A%22deploytime%22%2C%22type%22%3A%22desc%22%7D%5D&p="
	GetFirstUrl(url)
	url = "https://www.miit.gov.cn/search-front-server/api/search/info?websiteid=110000000000000&scope=basic&q=&pg=10&cateid=57&pos=title_text%2Cinfocontent%2Ctitlepy&selectFields=title,url,&group=distinct&level=6&sortFields=%5B%7B%22name%22%3A%22deploytime%22%2C%22type%22%3A%22desc%22%7D%5D&p=0"
	GetDetailPageUrl(url, "")
	url = "https://www.miit.gov.cn/zwgk/zcwj/wjfb/qt/art/2021/art_9b8e1e5711c54cec9ca0cda73ac36c40.html"
	url = "https://www.miit.gov.cn/zwgk/zcwj/wjfb/tg/art/2021/art_b5fbf1361d314615a63f8dae9646c511.html"
	GetHtmlInfo(url, "https://www.miit.gov.cn")

}


