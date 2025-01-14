package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/team/esgrammar/parser"
)

func main() {
	// 示例 IQL 查询
	iql := "(COMMENT=\"red1\" OR 'color1'=\"blue\") and ('pric1e'=2000 and 'mode1l'=\"hyundai\") and '计划部署时间' > sum(1,100) and '类型' in [ \"线上Bug\"] and '创建时间' >= '2024-01-01' and '创建时间' <= '2024-03-31' and '所属产品' = \"code\" and 'xxx' not in [\"a\",\"b\",\"c\"]  and 'QA引入原因分析' in ( \"无法复现\") and '状态' != \"已取消\" order by \"orderField\" desc, \"orderField2\" desc,\"orderField3\" "

	// 创建词法分析器
	input := antlr.NewInputStream(iql)
	lexer := parser.NewIqlLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)

	// 创建解析器
	p := parser.NewIqlParser(tokens)

	// 解析并生成语法树
	tree := p.Query()

	// 创建自定义监听器
	listener := parser.NewCustomIqlListener()

	// 使用 ParseTreeWalker 遍历解析树
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	// 获取生成的 QueryDSL
	queryDSL := listener.GetQueryDSL()
	fmt.Println("Generated Query DSL:", queryDSL)

	// 输出所有字段
	fmt.Println("Fields:")
	for field := range listener.GetFields() {
		fmt.Println(field)
	}

	// sort
	fmt.Println("Sort:")
	sortMaps := listener.GetSortMaps()
	for index := range sortMaps {
		for field, order := range sortMaps[index] {
			fmt.Println(field, order)
		}
	}

}
