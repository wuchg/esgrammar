// Code generated from parser/Iql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Iql

import "github.com/antlr4-go/antlr/v4"

// IqlListener is a complete listener for a parse tree produced by IqlParser.
type IqlListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterIql_stmt is called when entering the iql_stmt production.
	EnterIql_stmt(c *Iql_stmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterOrdering_term is called when entering the ordering_term production.
	EnterOrdering_term(c *Ordering_termContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterLiteral_value is called when entering the literal_value production.
	EnterLiteral_value(c *Literal_valueContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterFunction_param is called when entering the function_param production.
	EnterFunction_param(c *Function_paramContext)

	// EnterLiteral_list is called when entering the literal_list production.
	EnterLiteral_list(c *Literal_listContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterState_name is called when entering the state_name production.
	EnterState_name(c *State_nameContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterCompare_dates is called when entering the compare_dates production.
	EnterCompare_dates(c *Compare_datesContext)

	// EnterDates is called when entering the dates production.
	EnterDates(c *DatesContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitIql_stmt is called when exiting the iql_stmt production.
	ExitIql_stmt(c *Iql_stmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitOrdering_term is called when exiting the ordering_term production.
	ExitOrdering_term(c *Ordering_termContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitLiteral_value is called when exiting the literal_value production.
	ExitLiteral_value(c *Literal_valueContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitFunction_param is called when exiting the function_param production.
	ExitFunction_param(c *Function_paramContext)

	// ExitLiteral_list is called when exiting the literal_list production.
	ExitLiteral_list(c *Literal_listContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitState_name is called when exiting the state_name production.
	ExitState_name(c *State_nameContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitCompare_dates is called when exiting the compare_dates production.
	ExitCompare_dates(c *Compare_datesContext)

	// ExitDates is called when exiting the dates production.
	ExitDates(c *DatesContext)
}
