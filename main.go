package main

import (
	"TouTiao_Code/active"
	"fmt"
	"net/http"
)

/*
	实现分页显示的新闻头条内容展示：
		1、

*/
func main() {
	//设置静态文件资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", active.ToIndex)
	http.HandleFunc("/news.html", active.News)
	//就是通过虚拟路径执行一样的查询新闻的函数
	http.HandleFunc("/news", active.News)

	http.HandleFunc("/index.html", active.ToIndex)

	err := http.ListenAndServe(":8066", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("服务器启动成功")
}
