// Code generated from soql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // soql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by soqlParser.
type soqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by soqlParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by soqlParser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by soqlParser#fieldList.
	VisitFieldList(ctx *FieldListContext) interface{}

	// Visit a parse tree produced by soqlParser#selectField.
	VisitSelectField(ctx *SelectFieldContext) interface{}

	// Visit a parse tree produced by soqlParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by soqlParser#filterScope.
	VisitFilterScope(ctx *FilterScopeContext) interface{}

	// Visit a parse tree produced by soqlParser#SoqlFieldReference.
	VisitSoqlFieldReference(ctx *SoqlFieldReferenceContext) interface{}

	// Visit a parse tree produced by soqlParser#SoqlFunctionCall.
	VisitSoqlFunctionCall(ctx *SoqlFunctionCallContext) interface{}

	// Visit a parse tree produced by soqlParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by soqlParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by soqlParser#whereFields.
	VisitWhereFields(ctx *WhereFieldsContext) interface{}

	// Visit a parse tree produced by soqlParser#whereField.
	VisitWhereField(ctx *WhereFieldContext) interface{}

	// Visit a parse tree produced by soqlParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by soqlParser#orderClause.
	VisitOrderClause(ctx *OrderClauseContext) interface{}

	// Visit a parse tree produced by soqlParser#soqlValue.
	VisitSoqlValue(ctx *SoqlValueContext) interface{}

	// Visit a parse tree produced by soqlParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by soqlParser#soqlFilteringExpression.
	VisitSoqlFilteringExpression(ctx *SoqlFilteringExpressionContext) interface{}

	// Visit a parse tree produced by soqlParser#groupClause.
	VisitGroupClause(ctx *GroupClauseContext) interface{}

	// Visit a parse tree produced by soqlParser#havingConditionExpression.
	VisitHavingConditionExpression(ctx *HavingConditionExpressionContext) interface{}

	// Visit a parse tree produced by soqlParser#offsetClause.
	VisitOffsetClause(ctx *OffsetClauseContext) interface{}

	// Visit a parse tree produced by soqlParser#viewClause.
	VisitViewClause(ctx *ViewClauseContext) interface{}

	// Visit a parse tree produced by soqlParser#soslLiteral.
	VisitSoslLiteral(ctx *SoslLiteralContext) interface{}

	// Visit a parse tree produced by soqlParser#soslQuery.
	VisitSoslQuery(ctx *SoslQueryContext) interface{}

	// Visit a parse tree produced by soqlParser#soslReturningObject.
	VisitSoslReturningObject(ctx *SoslReturningObjectContext) interface{}

	// Visit a parse tree produced by soqlParser#apexIdentifier.
	VisitApexIdentifier(ctx *ApexIdentifierContext) interface{}

	// Visit a parse tree produced by soqlParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
