//go:generate ginger $GOFILE
package main

//@ginger
type TaipeiCompany struct {
	Data struct {
		代表人姓名    string     `json:"代表人姓名"`
		公司名稱     string     `json:"公司名稱"`
		公司所在地    string     `json:"公司所在地"`
		公司狀況     string     `json:"公司狀況"`
		實收資本額_元_ string     `json:"實收資本額(元)"`
		所營事業資料   [][]string `json:"所營事業資料"`
		最後核准變更日期 struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"最後核准變更日期"`
		核准設立日期 struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"核准設立日期"`
		登記機關  string        `json:"登記機關"`
		經理人名單 []interface{} `json:"經理人名單"`
		股權狀況  string        `json:"股權狀況"`
		董監事名單 []struct {
			出資額   string `json:"出資額"`
			姓名    string `json:"姓名"`
			序號    string `json:"序號"`
			所代表法人 string `json:"所代表法人"`
			職稱    string `json:"職稱"`
		} `json:"董監事名單"`
		資本總額_元_ string `json:"資本總額(元)"`
	} `json:"data"`
}
