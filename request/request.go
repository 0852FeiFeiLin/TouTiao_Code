package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)
//发起请求，返回响应
func Request(method string,url string,body io.Reader) (re []byte){
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("客户端发起请求失败")
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if response.StatusCode != 200{
		fmt.Println(err.Error())
		return
	}
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
	return bytes
}