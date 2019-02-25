package parser

import (
	"github.com/sausheong/gwp/crawler/engine"
	"regexp"
)

const profileRe  = `<div data-v-5b109fc3="" class="des f-cl">([^<])*</div>`

/**
解析个人信息页面，得到个人数据
 */
func ParseProfile(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(profileRe)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "UserInfo " + string(m[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
