package entity
//详细数据
type Data struct {
	//新闻id
	Uniquekey string `json:"uniquekey"`
	//标题
	Title string `json:"title"`
	//日期
	Date string `json:"date"`
	//新闻分类
	Ccategory string `json:"ccategory"`
	//新闻来源
	AuthorName string `json:"author_name"`
	//url地址
	Url string `json:"url"`
	//新闻图片链接
	Thumbnail_pic_s string  `json:"thumbnail_pic_s"`
	Thumbnail_pic_s02 string `json:"thumbnail_pic_s_02"`
	Tthumbnail_pic_s03 string `json:"tthumbnail_pic_s_03"`
	//是否有新闻内容
	Is_content string `json:"is_content"`

}