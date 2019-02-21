package main

import (
	"github.com/sausheong/gwp/crawler/engine"
	"github.com/sausheong/gwp/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
