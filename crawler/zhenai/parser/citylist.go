package parser

import (
	"github.com/sausheong/gwp/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

/**
解析主页面，得到城市的URL和城市名称
 */
func ParseCityList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityListRe)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 2
	for index, m := range matches {
		result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		if index == limit{
			break
		}
	}
	return result
}
