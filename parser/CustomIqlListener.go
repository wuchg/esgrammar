package parser

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

type (
	Stack struct {
		top    *node
		length int
		lock   *sync.RWMutex
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Push(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	n := &node{value, s.top}
	s.top = n
	s.length++
}

// CustomIqlListener 实现了 ANTLR 监听器接口，处理 IQL 表达式的解析
type CustomIqlListener struct {
	*BaseIqlListener
	query           strings.Builder
	registeredFuncs map[string]func([]interface{}) interface{}
	funcResults     map[string]string
	operatorMap     map[string]string
	sortMaps        []map[string]string
	fields          map[string]struct{}
	contextStack    *Stack
}

// NewCustomIqlListener 创建一个新的监听器实例
func NewCustomIqlListener() *CustomIqlListener {
	listener := &CustomIqlListener{
		registeredFuncs: make(map[string]func([]interface{}) interface{}),
		funcResults:     make(map[string]string),
		operatorMap:     make(map[string]string),
		sortMaps:        make([]map[string]string, 10),
		fields:          make(map[string]struct{}),
		contextStack:    NewStack(),
	}

	listener.registerBuiltInFunctions()

	// 初始化操作符映射
	listener.operatorMap["="] = "{\"term\": { \"#FIELD_KEY#\" :\"#FIELD_VALUE#\" }}"
	listener.operatorMap["!="] = "{\"bool\": {\"must_not\": { \"term\": {\"#FIELD_KEY#\" :\"#FIELD_VALUE#\"}}}}"
	listener.operatorMap["~"] = "{\"match\": {\"#FIELD_KEY#\" :\"#FIELD_VALUE#\"}}"
	listener.operatorMap["!~"] = "{\"bool\": {\"must_not\": { \"match\": {\"#FIELD_KEY#\" :\"#FIELD_VALUE#\"}}}}"
	listener.operatorMap["<"] = "{\"range\": { \"#FIELD_KEY#\": { \"lt\" : \"#FIELD_VALUE#\" }}}"
	listener.operatorMap["<="] = "{\"range\": { \"#FIELD_KEY#\": { \"lte\" : \"#FIELD_VALUE#\" }}}"
	listener.operatorMap[">"] = "{\"range\": { \"#FIELD_KEY#\": { \"gt\" : \"#FIELD_VALUE#\" }}}"
	listener.operatorMap[">="] = "{\"range\": { \"#FIELD_KEY#\": { \"gte\" : \"#FIELD_VALUE#\" }}}"
	listener.operatorMap["IN"] = "{\"terms\": { \"#FIELD_KEY#\" : [ #FIELD_VALUE# ] }}"
	listener.operatorMap["NOTIN"] = "{\"bool\": {\"must_not\": { \"terms\": { \"#FIELD_KEY#\" : [ #FIELD_VALUE# ] }}}}"

	return listener
}

// 注册内置函数
func (listener *CustomIqlListener) registerBuiltInFunctions() {
	// 在这里注册内置函数
	listener.registeredFuncs["sum"] = func(params []interface{}) interface{} {
		var sum float64
		for _, param := range params {
			fmt.Printf("param: %v\n", param)
			p, error := strconv.ParseFloat(param.(string), 32)
			if error != nil {
				fmt.Printf("error: %v\n", error)
			}
			sum += p

		}
		fmt.Printf("sum: %v\n", sum)
		return sum
	}
}

// 进入查询时
func (listener *CustomIqlListener) EnterQuery(ctx *QueryContext) {
	listener.query.WriteString("{ \"query\": ")
}

// 离开查询时
func (listener *CustomIqlListener) ExitQuery(ctx *QueryContext) {
	listener.query.WriteString("}")
}

// 进入表达式时
func (listener *CustomIqlListener) EnterExpr(ctx *ExprContext) {
	listener.contextStack.Push(ctx)
	// 处理表达式，判断是否是括号中的表达式
	if ctx.OPEN_PAR() != nil {
		innerCtx := ctx.AllExpr()[0]
		if innerCtx.K_AND() != nil {
			listener.query.WriteString("{ \"bool\": { \"must\": [")
		} else if innerCtx.K_OR() != nil {
			listener.query.WriteString("{ \"bool\": { \"should\": [")
		} else if innerCtx.AllK_NOT() != nil && len(innerCtx.AllK_NOT()) > 0 {
			listener.query.WriteString("{ \"bool\": { \"must_not\": [")
		}
	} else {
		// 普通表达式
		if listener.contextStack.Len() == 1 {
			if ctx.K_AND() != nil {
				listener.query.WriteString("{ \"bool\": { \"must\": [")
			} else if ctx.K_OR() != nil {
				listener.query.WriteString("{ \"bool\": { \"should\": [")
			} else if ctx.AllK_NOT() != nil && len(ctx.AllK_NOT()) > 0 {
				listener.query.WriteString("{ \"bool\": { \"must_not\": [")
			}
		}
	}
}

// 离开表达式时
func (listener *CustomIqlListener) ExitExpr(ctx *ExprContext) {

	listener.contextStack.Pop()
	if ctx.K_AND() != nil || ctx.K_OR() != nil || (ctx.AllK_NOT() != nil && len(ctx.AllK_NOT()) > 0) {
		if listener.isNotLast(ctx) {
			listener.query.WriteString(",")
		} else {
			listener.query.WriteString("]}}")
		}
		return
	}

	if ctx.Operator() == nil && ctx.CLOSE_PAR() != nil {
		if listener.isNotLast(ctx) {
			listener.query.WriteString(",")
		}
		return
	}
	// 处理字段和值，并替换操作符
	operator := strings.ToUpper(ctx.Operator().GetText())
	var field string = ""
	if ctx.Field() != nil {
		field = ctx.Field().GetText()
	} else if ctx.STRING_LITERAL() != nil {
		field = ctx.STRING_LITERAL().GetText()
	} else {
		fmt.Printf("field is empty")
	}
	field = strings.ReplaceAll(field, "\"", "")
	field = strings.ReplaceAll(field, "'", "")
	if field == "" {
		fmt.Printf("field is empty")
	}
	listener.fields[field] = struct{}{}
	var value string = ""
	if ctx.Literal_value() != nil {
		value = ctx.Literal_value().GetText()
		value = strings.ReplaceAll(value, "\"", "")
		value = strings.ReplaceAll(value, "'", "")
	} else if ctx.Literal_list() != nil {
		value = strings.ReplaceAll(ctx.Literal_list().GetText(), "[", "")
		value = strings.ReplaceAll(value, "]", "")
		value = strings.ReplaceAll(value, "(", "")
		value = strings.ReplaceAll(value, ")", "")
	} else if ctx.Function() != nil {
		value = ctx.Function().GetText()
		fmt.Printf("Function value: %v\n", value)
		value = listener.funcResults[value]
	} else {
		fmt.Printf("value is empty")
	}

	// 使用 operatorMap 替换占位符
	listener.query.WriteString(strings.Replace(strings.Replace(listener.operatorMap[operator], "#FIELD_KEY#", field, 1), "#FIELD_VALUE#", value, 1))
	if listener.isNotLast(ctx) {
		listener.query.WriteString(",")
	}
}

func (s *CustomIqlListener) ExitFunction(ctx *FunctionContext) {
	functionName := ctx.Function_name().GetText()
	params := make([]interface{}, 0)
	for _, param := range ctx.AllFunction_param() {
		if param.Literal_value() != nil {
			params = append(params, param.Literal_value().GetText())
		}
	}

	result := s.registeredFuncs[functionName](params)

	stringParams := make([]string, len(params))
	for i, param := range params {
		stringParams[i] = param.(string)
	}
	functionSign := fmt.Sprintf("%s(%s)", functionName, strings.Join(stringParams, ","))
	fmt.Printf("functionName: %s, params: %v, result: %v\n", functionName, params, result)
	fmt.Printf("functionSign: %s\n", functionSign)

	s.funcResults[functionSign] = fmt.Sprintf("%v", result)
}

func (s *CustomIqlListener) ExitOrdering_term(ctx *Ordering_termContext) {

	if ctx.K_ORDER() == nil || ctx.K_BY() == nil {
		return
	}

	childCount := ctx.GetChildCount()
	var field string = ""
	var order string = "asc"

	// for i, child := range ctx.GetChildren() {
	// 	fmt.Printf("child: %d, %v\n", i, child)
	// }

	//skip K_ORDER and K_BY
	for i := 2; i < childCount; i++ {
		field = strings.ReplaceAll(ctx.GetChild(i).(antlr.ParseTree).GetText(), "\"", "")
		if i+1 < childCount {
			if ctx.GetChild(i+1).(antlr.ParseTree).GetText() == "," {
				order = "asc" // default order is asc
				i++
			} else {
				order = ctx.GetChild(i + 1).(antlr.ParseTree).GetText()
				i += 2
			}
		} else {
			order = "asc"
		}

		orderMap := make(map[string]string)
		orderMap[field] = order
		s.sortMaps = append(s.sortMaps, orderMap)
	}

}

func (s *CustomIqlListener) GetSortMaps() []map[string]string {
	return s.sortMaps
}

func (l *CustomIqlListener) isNotLast(ctx *ExprContext) bool {
	if l.contextStack.Len() > 0 {
		parentCtx, ok := l.contextStack.Peek().(*ExprContext)
		if ok {
			return len(parentCtx.AllExpr()) > 1 && parentCtx.AllExpr()[len(parentCtx.AllExpr())-1] != ctx
		}
	}
	return false
}

// EnterEveryRule is called when any rule is entered.
func (s *CustomIqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *CustomIqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// VisitTerminal is called when a terminal node is visited.
func (s *CustomIqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *CustomIqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// 获取生成的 QueryDSL
func (listener *CustomIqlListener) GetQueryDSL() string {
	return listener.query.String()
}

// 获取所有字段
func (listener *CustomIqlListener) GetFields() map[string]struct{} {
	return listener.fields
}
