package main

type TaipeiCity struct {
	Ginger_Created int32  `json:"ginger_created"`
	Ginger_Id      int32  `json:"ginger_id" gorm:"primary_key"`
	Address        string `json:"address"`
	Exit           string `json:"exit"`
}
