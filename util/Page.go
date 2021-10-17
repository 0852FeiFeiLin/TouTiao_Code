package util

import "strconv"

type Page struct {
	//上一页
	PrePage int64
	//下一页
	NextPage int64
	//当前页数  ==> pages
	CurrentPage int64
	//总页数
	PageCount int64
	//页数集合 [1,2,3,..n]
	PageNum []int
	//类型
	NewsType string
}
/*
 * 分页
 * tableName 数据体
 * count 数据总条数
 * currentPage 当前页数
 * pageSize 一页显示多少条数据
 */

//分页：传入三个参数，类型，总页数，当前页数
func (p *Page)ShowPage(NewsType string,PageCount int64,CurrentPage int64 ) string{
	//结果
	var res = ""
	//当前页数i == CurrentPage
	var i int64
	//把总页数转为int64 ===> string，方便后面的结果展示
	pageCount := strconv.FormatInt(PageCount, 10)//转换为10进制

	//判断返回页数是否大于5页
	if PageCount > 5{ //50 >5
		//1、页数大于等于46情况
		if PageCount -CurrentPage<5{ //50-46<5
			//避免超过总页数
			x := PageCount - CurrentPage -5 //50-46-5 = -1
			for i = x + CurrentPage +1;i<=PageCount;i++ { //-1+46+1= 46，47，48,49,50  循环5次
				//当前页数  int64 ==> string
				currentPage := strconv.FormatInt(i, 10)
				//结果展示
				res += `<li><a href="`  + `/news?type=` + NewsType + `&page=` + currentPage + `">` + currentPage + `</a></li>`
			}
			//返回结果string类型，然后再那边调用模板语法转换为html格式
			return res
		}
		//2、正常返回页数小于46情况
		for i = CurrentPage; i < CurrentPage +5;i++{ //循环五次
			//当前页数：int64 ==> string
			currentPage := strconv.FormatInt(i, 10)
			//结果
			res += `<li><a href="`  + `/news?type=` + NewsType + `&page=` + currentPage + `">` + currentPage + `</a></li>`
			//45,46,47,48,49
		}
		//加上总页数
		res += `<li><a href="`  + `/news?type=` + NewsType + `&page=` + pageCount + `">` + pageCount + `</a></li>`
		return res
		//返回页数小于五条的情况
	}else if PageCount <= 5 && PageCount != 0 {
		for i = 1;i<PageCount;i++ {
			currentPage := strconv.FormatInt(i, 10)
			//结果
			res += `<li><a href="`  + `/news?type=` + NewsType + `&page=` + currentPage + `">` + currentPage + `</a></li>`
		}
			return res
	}
	//返回空
return ""
}


/*
func (p *Page)Paginaction(info entity.Info,count int64,currentPage int64,pageSize int)([]entity.Info,Page, int64, error){
												//返回两部分数据，一个是查询到的后端结构体数据，一个是显示给前端的
	//总页数  = 数据总条数 / 一页展示数据
	PageCount := math.Ceil(float64(count) / float64(pageSize))
	//分页操作
	//news ,err := service.LimitPage(info,currentPage-1, pageSize,PageCount)//返回一个结构体数据，
	//就是把数据源和当前页数，和一页展示多少数据和总页数传递给分页方法
	fmt.Println(currentPage,pageSize,PageCount)

	//接下里就要给这个结构体里面的属性赋值，因为后边是要穿给别的方法的，前端展示也依赖这个结构体
	page.PageCount = int64(PageCount)// 等操作
	//return 返回查询到的后端结构体数据news + 响应在前端的前端展示数据data + 总页数pagecount + error

}
*/