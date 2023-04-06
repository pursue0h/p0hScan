package get

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// 测试用
//var url = "http://pursue0h.top"

// 发送请求
func Requset(url string) *http.Response {
	response, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	return response
}

// 获取响应的Header
func Request_header(Response *http.Response) http.Header {
	header := Response.Header
	return header
}

// header转字符串
func Header_string(header http.Header) string {
	x := ""
	for key, value := range header {
		x = x + " " + key
		for i := 0; i < len(value); i++ {
			x = x + " " + value[i]
		}
	}
	return x
}

// 获取响应的Header中的server
func Request_Server(response *http.Response) string {
	server := response.Header["Server"][0] //在这里用这种格式进行匹配是因为Header是map[string][]string结构，匹配了键值对
	return server
}

// 获取响应的Body
func Requset_Body(response *http.Response) []byte {
	response_body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return response_body
}

//获取响应的TLS证书 (未实现)
/*
func Request_TLS(response *http.Response){
	response_TLS,err3:=
}
*/

// 获取Body中的title
func Body_title(response_body []byte) string {
	reg_title := regexp.MustCompile("<title>(.*?)</title>")
	title := reg_title.FindStringSubmatch(string(response_body))[1]
	return title
}

// 获取Body中的banner(这个banner在好多网页中都没有，在编写进行测试的时候可以自己在服务器html中插入一些<banner>测试一下)
func Body_banner(response_body []byte) [][]string {
	reg_banner := regexp.MustCompile(`(?im)<\s*banner.*>(.*?)<\s*/\s*banner>`)
	banner := reg_banner.FindAllStringSubmatch(string(response_body), -1)
	return banner
}

/*以下代码为了优化将其删除

// 提供url返回body，方便外部引用
func Finall_get_body(url string) []byte {
	finall_body := Requset_Body(Requset(url))
	return finall_body
}

// 提供url返回title，方便外部引用
func Finall_get_title(url string) string {
	finall_title := Body_title(Requset_Body(Requset(url)))
	return finall_title
}

// 提供url返回server，方便外部引用
func Finall_get_server(url string) string {
	finall_server := Request_Server(Requset(url))
	return finall_server
}

// 提供url返回banner，方便外部引用
func Finall_get_banner(url string) [][]string {
	finall_banner := Body_banner(Requset_Body(Requset(url)))
	return finall_banner
}
*/
// 测试用
/*
func main() {

	fmt.Println(Body_banner(Requset_Body(Requset(url))))
}
*/
