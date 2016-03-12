package service

import (
	//"code.google.com/p/mahonia"
	//"fmt"
	"mysql"
	"regexp"
	"strings"
	"time"
	"utils"
)

func Crawler() {

	for _, val := range SourceUrls {
		url := val.Url
		//deep := val.Deep

		src := utils.HttpGet(url)

		//将HTML标签全转换成小写
		re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
		src = re.ReplaceAllStringFunc(src, strings.ToLower)

		//去除STYLE
		re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
		src = re.ReplaceAllString(src, "")

		//去除SCRIPT
		re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
		src = re.ReplaceAllString(src, "")

		//去除所有尖括号内的HTML代码，并换成换行符
		re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
		src = re.ReplaceAllString(src, "\n")

		//去除连续的换行符
		re, _ = regexp.Compile("\\s{2,}")
		src = re.ReplaceAllString(src, "\n")
		src = strings.TrimSpace(src)

		src = strings.Replace(src, "\r", " ", -1)
		src = strings.Replace(src, "\n", " ", -1)
		mysql.Insert("insert into source_data (url,data,update_time) values (?,?,?)", url, src, time.Now().Unix())
	}

}
