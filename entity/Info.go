package entity
//大的内容展示区
type Info struct {
	//返回说明
	Reason string `json:"reason"`
	//结果集
	Result Result `json:"result"`
	//返回码 0代表成功
	ErrorCode int `json:"error_code"`
}
