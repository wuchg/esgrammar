// Code generated from parser/Iql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Iql

import "github.com/antlr4-go/antlr/v4"

// BaseIqlListener is a complete listener for a parse tree produced by IqlParser.
type BaseIqlListener struct{}

var _ IqlListener = &BaseIqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseIqlListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseIqlListener) ExitQuery(ctx *QueryContext) {}

// EnterIql_stmt is called when production iql_stmt is entered.
func (s *BaseIqlListener) EnterIql_stmt(ctx *Iql_stmtContext) {}

// ExitIql_stmt is called when production iql_stmt is exited.
func (s *BaseIqlListener) ExitIql_stmt(ctx *Iql_stmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseIqlListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseIqlListener) ExitExpr(ctx *ExprContext) {}

// EnterOrdering_term is called when production ordering_term is entered.
func (s *BaseIqlListener) EnterOrdering_term(ctx *Ordering_termContext) {}

// ExitOrdering_term is called when production ordering_term is exited.
func (s *BaseIqlListener) ExitOrdering_term(ctx *Ordering_termContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseIqlListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseIqlListener) ExitOperator(ctx *OperatorContext) {}

// EnterLiteral_value is called when production literal_value is entered.
func (s *BaseIqlListener) EnterLiteral_value(ctx *Literal_valueContext) {}

// ExitLiteral_value is called when production literal_value is exited.
func (s *BaseIqlListener) ExitLiteral_value(ctx *Literal_valueContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseIqlListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseIqlListener) ExitFunction(ctx *FunctionContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseIqlListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseIqlListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterFunction_param is called when production function_param is entered.
func (s *BaseIqlListener) EnterFunction_param(ctx *Function_paramContext) {}

// ExitFunction_param is called when production function_param is exited.
func (s *BaseIqlListener) ExitFunction_param(ctx *Function_paramContext) {}

// EnterLiteral_list is called when production literal_list is entered.
func (s *BaseIqlListener) EnterLiteral_list(ctx *Literal_listContext) {}

// ExitLiteral_list is called when production literal_list is exited.
func (s *BaseIqlListener) ExitLiteral_list(ctx *Literal_listContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseIqlListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseIqlListener) ExitKeyword(ctx *KeywordContext) {}

// EnterState_name is called when production state_name is entered.
func (s *BaseIqlListener) EnterState_name(ctx *State_nameContext) {}

// ExitState_name is called when production state_name is exited.
func (s *BaseIqlListener) ExitState_name(ctx *State_nameContext) {}

// EnterField is called when production field is entered.
func (s *BaseIqlListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseIqlListener) ExitField(ctx *FieldContext) {}

// EnterCompare_dates is called when production compare_dates is entered.
func (s *BaseIqlListener) EnterCompare_dates(ctx *Compare_datesContext) {}

// ExitCompare_dates is called when production compare_dates is exited.
func (s *BaseIqlListener) ExitCompare_dates(ctx *Compare_datesContext) {}

// EnterDates is called when production dates is entered.
func (s *BaseIqlListener) EnterDates(ctx *DatesContext) {}

// ExitDates is called when production dates is exited.
func (s *BaseIqlListener) ExitDates(ctx *DatesContext) {}
