package json

import (
	"fmt"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go" //看的o2的，本来想用fastjson的，但是好像有些问题，所以就换成这个了 更加灵活一些
)

// 定义fofa的结构体 来源：https://mholt.github.io/json-to-go/  json转struct
type json_struct []struct {
	RuleID         string `json:"rule_id"`
	Level          string `json:"level"`
	Softhard       string `json:"softhard"`
	Product        string `json:"product"`
	Company        string `json:"company"`
	Category       string `json:"category"`
	ParentCategory string `json:"parent_category"`
	Rules          [][]struct {
		Match   string `json:"match"`
		Content string `json:"content"`
	} `json:"rules"`
}

// ConfigCompatibleWithStandardLibrary尝试与标准库行为100%兼容配置
var json = jsoniter.ConfigCompatibleWithStandardLibrary

var Fofa_data json_struct

// 读取json并存到struct中
func Read_json() json_struct {

	//打开文件
	fofa, err1 := os.Open("./json/fofa.json")
	if err1 != nil {
		fmt.Println(err1)
	}
	//关闭文件
	defer fofa.Close()
	//存入struct中
	fofa_read, _ := ioutil.ReadAll(fofa)
	err2 := json.Unmarshal(fofa_read, &Fofa_data)
	if err2 != nil {
		fmt.Println(err2)
	}
	return Fofa_data
}

// 测试用
/*
func main() {
	fmt.Println(Read_json())

}
*/
