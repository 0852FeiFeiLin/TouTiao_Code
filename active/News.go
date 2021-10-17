package active

import (
	"TouTiao_Code/page_parse"
	"TouTiao_Code/request"
	"TouTiao_Code/util"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)
//展示新闻详情
func News(w http.ResponseWriter ,r *http.Request){
	//通过该方法找到url?后面的值
	value := r.URL.Query();
	fmt.Println(value)
	//根据具体的属性判断url里面的type和page
	vtype := value.Get("type")
	pages := value.Get("page")

	//设置默认值
	if vtype == ""{
		vtype = "top"
		//如果为空就是top(默认类型)
	}
	if pages == ""{
		 pages   = "1"
	}
	fmt.Println(vtype,pages)

	//先把pages当前页数进行处理
	//实例化后面要传输的结构体
	var p util.Page
	//string ===> int64
	pagesInt, _ := strconv.ParseInt(pages, 10, 64)
	//Page结构体 赋值
	p.CurrentPage = pagesInt  //当前页数
	p.NewsType = vtype    //新闻类型
	p.PrePage = pagesInt - 1    //上一页
	p.NextPage = pagesInt +1    //下一页
	p.PageCount = 50    //总页数

	//拼接URL
	urlPATH := APIURL + "?type=" + vtype + "&page=" + pages + "&page_size=&is_filter=&key=" + KEY
	fmt.Println(urlPATH)
	//发起请求，获得响应
	res := request.Request("POST", urlPATH, nil)
	//处理响应，反序列化
	info := page_parse.ParseJson(res)
	//打印响应数据
	//fmt.Println(info)
	//创建模板
	files, err := template.ParseFiles("./view/news.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//新闻详情
	news := info.Result.Data

	//将结构体数据转为切片，利用于模板语法,利用模板语法还能设置模板中的type和page
	data := map[string] interface{}{
		//news 就是新闻详情，
		"News" : news,
		//"Type" : vtype,
		//page里面传Page结构体，里面有类型，页数，上一页，下一页那些属性
		"Page" : p,
		//showPage() 就是动态生成页码的函数  str2html 就是string转为html
		"Pages" : str2Html(p.ShowPage(p.NewsType,p. PageCount,p.CurrentPage)),
		//                          类型         总页数           当前页数

	}
	//模板发送map过去，前端通过遍历key去
	files.Execute(w,data)
}
//string ==>html
func str2Html(content string)template.HTML{
	return template.HTML(content)
}