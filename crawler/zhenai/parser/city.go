package parser

import (
	"github.com/sausheong/gwp/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

/**
解析城市页面，获取用户的名称和URL
 */
func ParseCity(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityRe)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseProfile,
		})
	}
	return result
}
