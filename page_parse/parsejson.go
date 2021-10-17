package page_parse

import (
	"TouTiao_Code/entity"
	"encoding/json"
	"fmt"
)

func ParseJson(news []byte) (re entity.Info){
	var info entity.Info
	jsonerr := json.Unmarshal(news,&info)
	if jsonerr != nil {
		fmt.Println(jsonerr.Error())
		fmt.Println("json解析出现错误")
		return
	}
	return info
}
