# TouTiao_Code
这是一个新闻头条查询展示的案例


这是利用api查询到新闻详情，利用了分页查询和分页展示，分页展示在前端，然后动态生成页码，前端单行栏就是指定的type,
然后通过每一个a标签发起get请求，get请求的type和page从当前news.html中获取，然后再后端发起请求的方法中设置。默认的
type = top，page=1.
