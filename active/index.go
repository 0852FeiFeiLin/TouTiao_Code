package active

import (
	//"TouTiao_Code/service"
	"fmt"
	"html/template"
	"net/http"
)
const APIURL = "http://v.juhe.cn/toutiao/index"
const KEY = "60b8c5f36ef97f9ea371d3e58fd221a6"
//跳转到index页面
func ToIndex(w http.ResponseWriter,r *http.Request){
	files, err := template.ParseFiles("./view/index.html")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("跳转主页失败")
		return
	}
	files.Execute(w,nil)
}
//直接使用News方法就行了，这里重复了
//func Index(w http.ResponseWriter,r *http.Request){
//	//解析前端表单
//	err := r.ParseForm()
//	if err != nil {
//		fmt.Println(err.Error())
//		fmt.Println("解析表单数据失败")
//		return
//	}
//
//	vtype := r.FormValue("types")//根据name获取属性
//	page := r.FormValue("page")
//	pageSize := r.FormValue("page_size")
//	isFilter := r.FormValue("is_filter")
//
//	/*
//		对用户输入的数据进行检查，如果不符合已经有的数据，进行处理
//	*/
//
//	fmt.Println(vtype)
//	fmt.Println(vtype , page , pageSize , isFilter )
//	//http://v.juhe.cn/toutiao/index?top=yule&page=2&page_size=3&is_filter=1&key=7ade132af1a073b66f175798a7adf0c4
//	//定义url，发起请求
//
//	//拼接url
//	urlPATH := APIURL + "?type=" + vtype + "&page=" + page + "&page_size=" + pageSize + "&is_filter=" + isFilter + "&key=" + KEY
//	fmt.Println(urlPATH)
//	//发起请求,返回响应体
//	news := request.Request("POST", urlPATH, nil)
//	//反序列化
//	info := page_parse.ParseJson(news)
//	//fmt.Println(info)//请求到的数据
//
//	/*
//		注意:这里返回的数据都是一些新闻标题和摘要，还有超链接，重点：超链接，我们把超链接截取下来，供我们所用
//	*/
//
///
//	//创建模板，跳转到news页面里面
//	files, err := template.ParseFiles("./view/news.html")
//	data := map[string][] entity.Data{
//		"News" : info.Result.Data,
//	}
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//
//	//模板发送map，前端通过string类型的“News” key 遍历map
//	files.Execute(w,data)
//
//}

