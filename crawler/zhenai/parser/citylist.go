package parser

import (
	"fmt"
	"github.com/sausheong/gwp/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityListRe)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		fmt.Printf("City: %s URL: %s \n", m[2], m[1])
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
