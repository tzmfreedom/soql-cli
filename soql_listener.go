// Code generated from soql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // soql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// soqlListener is a complete listener for a parse tree produced by soqlParser.
type soqlListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterSelectField is called when entering the selectField production.
	EnterSelectField(c *SelectFieldContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterFilterScope is called when entering the filterScope production.
	EnterFilterScope(c *FilterScopeContext)

	// EnterSoqlFieldReference is called when entering the SoqlFieldReference production.
	EnterSoqlFieldReference(c *SoqlFieldReferenceContext)

	// EnterSoqlFunctionCall is called when entering the SoqlFunctionCall production.
	EnterSoqlFunctionCall(c *SoqlFunctionCallContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterWhereFields is called when entering the whereFields production.
	EnterWhereFields(c *WhereFieldsContext)

	// EnterWhereField is called when entering the whereField production.
	EnterWhereField(c *WhereFieldContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterOrderClause is called when entering the orderClause production.
	EnterOrderClause(c *OrderClauseContext)

	// EnterSoqlValue is called when entering the soqlValue production.
	EnterSoqlValue(c *SoqlValueContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterSoqlFilteringExpression is called when entering the soqlFilteringExpression production.
	EnterSoqlFilteringExpression(c *SoqlFilteringExpressionContext)

	// EnterGroupClause is called when entering the groupClause production.
	EnterGroupClause(c *GroupClauseContext)

	// EnterHavingConditionExpression is called when entering the havingConditionExpression production.
	EnterHavingConditionExpression(c *HavingConditionExpressionContext)

	// EnterOffsetClause is called when entering the offsetClause production.
	EnterOffsetClause(c *OffsetClauseContext)

	// EnterViewClause is called when entering the viewClause production.
	EnterViewClause(c *ViewClauseContext)

	// EnterSoslLiteral is called when entering the soslLiteral production.
	EnterSoslLiteral(c *SoslLiteralContext)

	// EnterSoslQuery is called when entering the soslQuery production.
	EnterSoslQuery(c *SoslQueryContext)

	// EnterSoslReturningObject is called when entering the soslReturningObject production.
	EnterSoslReturningObject(c *SoslReturningObjectContext)

	// EnterApexIdentifier is called when entering the apexIdentifier production.
	EnterApexIdentifier(c *ApexIdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitSelectField is called when exiting the selectField production.
	ExitSelectField(c *SelectFieldContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitFilterScope is called when exiting the filterScope production.
	ExitFilterScope(c *FilterScopeContext)

	// ExitSoqlFieldReference is called when exiting the SoqlFieldReference production.
	ExitSoqlFieldReference(c *SoqlFieldReferenceContext)

	// ExitSoqlFunctionCall is called when exiting the SoqlFunctionCall production.
	ExitSoqlFunctionCall(c *SoqlFunctionCallContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitWhereFields is called when exiting the whereFields production.
	ExitWhereFields(c *WhereFieldsContext)

	// ExitWhereField is called when exiting the whereField production.
	ExitWhereField(c *WhereFieldContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitOrderClause is called when exiting the orderClause production.
	ExitOrderClause(c *OrderClauseContext)

	// ExitSoqlValue is called when exiting the soqlValue production.
	ExitSoqlValue(c *SoqlValueContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitSoqlFilteringExpression is called when exiting the soqlFilteringExpression production.
	ExitSoqlFilteringExpression(c *SoqlFilteringExpressionContext)

	// ExitGroupClause is called when exiting the groupClause production.
	ExitGroupClause(c *GroupClauseContext)

	// ExitHavingConditionExpression is called when exiting the havingConditionExpression production.
	ExitHavingConditionExpression(c *HavingConditionExpressionContext)

	// ExitOffsetClause is called when exiting the offsetClause production.
	ExitOffsetClause(c *OffsetClauseContext)

	// ExitViewClause is called when exiting the viewClause production.
	ExitViewClause(c *ViewClauseContext)

	// ExitSoslLiteral is called when exiting the soslLiteral production.
	ExitSoslLiteral(c *SoslLiteralContext)

	// ExitSoslQuery is called when exiting the soslQuery production.
	ExitSoslQuery(c *SoslQueryContext)

	// ExitSoslReturningObject is called when exiting the soslReturningObject production.
	ExitSoslReturningObject(c *SoslReturningObjectContext)

	// ExitApexIdentifier is called when exiting the apexIdentifier production.
	ExitApexIdentifier(c *ApexIdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
