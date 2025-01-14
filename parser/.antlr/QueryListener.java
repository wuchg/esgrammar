// Generated from /Users/wuchg/esgrammar/parser/Query.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link QueryParser}.
 */
public interface QueryListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by the {@code SearchQuery}
	 * labeled alternative in {@link QueryParser#query}.
	 * @param ctx the parse tree
	 */
	void enterSearchQuery(QueryParser.SearchQueryContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SearchQuery}
	 * labeled alternative in {@link QueryParser#query}.
	 * @param ctx the parse tree
	 */
	void exitSearchQuery(QueryParser.SearchQueryContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CountQuery}
	 * labeled alternative in {@link QueryParser#query}.
	 * @param ctx the parse tree
	 */
	void enterCountQuery(QueryParser.CountQueryContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CountQuery}
	 * labeled alternative in {@link QueryParser#query}.
	 * @param ctx the parse tree
	 */
	void exitCountQuery(QueryParser.CountQueryContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#searchClause}.
	 * @param ctx the parse tree
	 */
	void enterSearchClause(QueryParser.SearchClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#searchClause}.
	 * @param ctx the parse tree
	 */
	void exitSearchClause(QueryParser.SearchClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#countClause}.
	 * @param ctx the parse tree
	 */
	void enterCountClause(QueryParser.CountClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#countClause}.
	 * @param ctx the parse tree
	 */
	void exitCountClause(QueryParser.CountClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#whereClause}.
	 * @param ctx the parse tree
	 */
	void enterWhereClause(QueryParser.WhereClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#whereClause}.
	 * @param ctx the parse tree
	 */
	void exitWhereClause(QueryParser.WhereClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#condition}.
	 * @param ctx the parse tree
	 */
	void enterCondition(QueryParser.ConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#condition}.
	 * @param ctx the parse tree
	 */
	void exitCondition(QueryParser.ConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#andCondition}.
	 * @param ctx the parse tree
	 */
	void enterAndCondition(QueryParser.AndConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#andCondition}.
	 * @param ctx the parse tree
	 */
	void exitAndCondition(QueryParser.AndConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#orCondition}.
	 * @param ctx the parse tree
	 */
	void enterOrCondition(QueryParser.OrConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#orCondition}.
	 * @param ctx the parse tree
	 */
	void exitOrCondition(QueryParser.OrConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#orderByClause}.
	 * @param ctx the parse tree
	 */
	void enterOrderByClause(QueryParser.OrderByClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#orderByClause}.
	 * @param ctx the parse tree
	 */
	void exitOrderByClause(QueryParser.OrderByClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#limitClause}.
	 * @param ctx the parse tree
	 */
	void enterLimitClause(QueryParser.LimitClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#limitClause}.
	 * @param ctx the parse tree
	 */
	void exitLimitClause(QueryParser.LimitClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#offsetClause}.
	 * @param ctx the parse tree
	 */
	void enterOffsetClause(QueryParser.OffsetClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#offsetClause}.
	 * @param ctx the parse tree
	 */
	void exitOffsetClause(QueryParser.OffsetClauseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ComparisonCondition}
	 * labeled alternative in {@link QueryParser#conditionPart}.
	 * @param ctx the parse tree
	 */
	void enterComparisonCondition(QueryParser.ComparisonConditionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ComparisonCondition}
	 * labeled alternative in {@link QueryParser#conditionPart}.
	 * @param ctx the parse tree
	 */
	void exitComparisonCondition(QueryParser.ComparisonConditionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code GroupCondition}
	 * labeled alternative in {@link QueryParser#conditionPart}.
	 * @param ctx the parse tree
	 */
	void enterGroupCondition(QueryParser.GroupConditionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code GroupCondition}
	 * labeled alternative in {@link QueryParser#conditionPart}.
	 * @param ctx the parse tree
	 */
	void exitGroupCondition(QueryParser.GroupConditionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code LikeCondition}
	 * labeled alternative in {@link QueryParser#conditionPart}.
	 * @param ctx the parse tree
	 */
	void enterLikeCondition(QueryParser.LikeConditionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code LikeCondition}
	 * labeled alternative in {@link QueryParser#conditionPart}.
	 * @param ctx the parse tree
	 */
	void exitLikeCondition(QueryParser.LikeConditionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IsNullCondition}
	 * labeled alternative in {@link QueryParser#conditionPart}.
	 * @param ctx the parse tree
	 */
	void enterIsNullCondition(QueryParser.IsNullConditionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IsNullCondition}
	 * labeled alternative in {@link QueryParser#conditionPart}.
	 * @param ctx the parse tree
	 */
	void exitIsNullCondition(QueryParser.IsNullConditionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code InCondition}
	 * labeled alternative in {@link QueryParser#conditionPart}.
	 * @param ctx the parse tree
	 */
	void enterInCondition(QueryParser.InConditionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code InCondition}
	 * labeled alternative in {@link QueryParser#conditionPart}.
	 * @param ctx the parse tree
	 */
	void exitInCondition(QueryParser.InConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#orderCondition}.
	 * @param ctx the parse tree
	 */
	void enterOrderCondition(QueryParser.OrderConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#orderCondition}.
	 * @param ctx the parse tree
	 */
	void exitOrderCondition(QueryParser.OrderConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#indexName}.
	 * @param ctx the parse tree
	 */
	void enterIndexName(QueryParser.IndexNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#indexName}.
	 * @param ctx the parse tree
	 */
	void exitIndexName(QueryParser.IndexNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#fieldName}.
	 * @param ctx the parse tree
	 */
	void enterFieldName(QueryParser.FieldNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#fieldName}.
	 * @param ctx the parse tree
	 */
	void exitFieldName(QueryParser.FieldNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(QueryParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(QueryParser.ValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link QueryParser#comparator}.
	 * @param ctx the parse tree
	 */
	void enterComparator(QueryParser.ComparatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link QueryParser#comparator}.
	 * @param ctx the parse tree
	 */
	void exitComparator(QueryParser.ComparatorContext ctx);
}