package match

import (
	"fmt"
	"strings"

	"p0hScan/get"
	"p0hScan/json"
)

//var url = "http://pursue0h.top" //测试用

func Match(url string) {
	//get包中获取需要匹配的东西

	response := get.Requset(url)
	response_header := get.Header_string(get.Request_header(response))
	response_server := get.Request_Server(response)
	response_body := get.Requset_Body(response)
	response_title := get.Body_title(response_body)
	response_banner := get.Body_banner(response_body)

	//解析一次获得struct
	Fofa_data := json.Read_json()

	//进行比对 使用for循环 i,j,k控制
	for i := 0; i < len(Fofa_data); i++ { //i控制struct的最外层

		for j := 0; j < len(Fofa_data[i].Rules); j++ { //j控制rules这一层

			//count作为rule中and的判断
			count := 0

			for k := 0; k < len(Fofa_data[i].Rules[j]); k++ { //k控制单个rule中多个contain
				switch Fofa_data[i].Rules[j][k].Match { //进行字符串比较
				case "header_contains":
					{
						if strings.Contains(string(response_header), Fofa_data[i].Rules[j][k].Content) {
							count++
						}
					}
				case "server_contains":
					{
						if response_server == Fofa_data[i].Rules[j][k].Content {
							count++
						}
					}
				case "body_contains":
					{
						if strings.Contains(string(response_body), Fofa_data[i].Rules[j][k].Content) {
							count++
						}
					}
				case "title_contains":
					{
						if response_title == Fofa_data[i].Rules[j][k].Content {
							count++
						}
					}
				case "banner_contains":
					{
						for m := 0; m < len(response_banner); m++ {
							if response_banner[m][1] == Fofa_data[i].Rules[j][k].Content {
								count++
								break
							}
						}
					}
				}
				if len(Fofa_data[i].Rules[j]) == count {
					fmt.Printf(">>>")
					fmt.Println(Fofa_data[i].Product)
				}
			}

		}
	}
}

// 测试用
/*
func main() {
	match(url)
}
*/
