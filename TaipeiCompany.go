//go:generate ginger $GOFILE
package main

//@ginger
type TaipeiCompany struct {
	Ginger_Created int32   `json:"ginger_created"`
	Ginger_Id      int32   `json:"ginger_id" gorm:"primary_key"`
	Data           Profile `json:"data"`
	DataId         int64
}

type Profile struct {
	Ceo     string `json:"代表人姓名"`
	Company string `json:"公司名稱"`
	Address string `json:"公司所在地"`
	Status  string `json:"公司狀況"`
	RealCap string `json:"實收資本額(元)"`
	//Metadata         [][]string `json:"所營事業資料"`
	LastApprovedDate struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	} `json:"最後核准變更日期"`
	ApprovedDate struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	} `json:"核准設立日期"`
	Gov string `json:"登記機關"`
	//ManagerList []interface{} `json:"經理人名單"`
	StockStatus string `json:"股權狀況"`
	BoardList   []struct {
		Invest   string `json:"出資額"`
		Name     string `json:"姓名"`
		SerialId string `json:"序號"`
		Rep      string `json:"所代表法人"`
		Title    string `json:"職稱"`
	} `json:"董監事名單"`
	TotalCap string `json:"資本總額(元)"`
}
