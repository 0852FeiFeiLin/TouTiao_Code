package entity

type  Result struct {
	//状态
	Stat string `json:"stat"`
	//内容
	Data []Data `json:"data"`
	//页数
	Page string `json:"page"`
	//条数
	PageSize string `json:"page_size"`

}
