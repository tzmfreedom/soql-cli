// Code generated from soql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // soql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesoqlListener is a complete listener for a parse tree produced by soqlParser.
type BasesoqlListener struct{}

var _ soqlListener = &BasesoqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesoqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesoqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesoqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesoqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BasesoqlListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BasesoqlListener) ExitQuery(ctx *QueryContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BasesoqlListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BasesoqlListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BasesoqlListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BasesoqlListener) ExitFieldList(ctx *FieldListContext) {}

// EnterSelectField is called when production selectField is entered.
func (s *BasesoqlListener) EnterSelectField(ctx *SelectFieldContext) {}

// ExitSelectField is called when production selectField is exited.
func (s *BasesoqlListener) ExitSelectField(ctx *SelectFieldContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BasesoqlListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BasesoqlListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterFilterScope is called when production filterScope is entered.
func (s *BasesoqlListener) EnterFilterScope(ctx *FilterScopeContext) {}

// ExitFilterScope is called when production filterScope is exited.
func (s *BasesoqlListener) ExitFilterScope(ctx *FilterScopeContext) {}

// EnterSoqlFieldReference is called when production SoqlFieldReference is entered.
func (s *BasesoqlListener) EnterSoqlFieldReference(ctx *SoqlFieldReferenceContext) {}

// ExitSoqlFieldReference is called when production SoqlFieldReference is exited.
func (s *BasesoqlListener) ExitSoqlFieldReference(ctx *SoqlFieldReferenceContext) {}

// EnterSoqlFunctionCall is called when production SoqlFunctionCall is entered.
func (s *BasesoqlListener) EnterSoqlFunctionCall(ctx *SoqlFunctionCallContext) {}

// ExitSoqlFunctionCall is called when production SoqlFunctionCall is exited.
func (s *BasesoqlListener) ExitSoqlFunctionCall(ctx *SoqlFunctionCallContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BasesoqlListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BasesoqlListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BasesoqlListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BasesoqlListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterWhereFields is called when production whereFields is entered.
func (s *BasesoqlListener) EnterWhereFields(ctx *WhereFieldsContext) {}

// ExitWhereFields is called when production whereFields is exited.
func (s *BasesoqlListener) ExitWhereFields(ctx *WhereFieldsContext) {}

// EnterWhereField is called when production whereField is entered.
func (s *BasesoqlListener) EnterWhereField(ctx *WhereFieldContext) {}

// ExitWhereField is called when production whereField is exited.
func (s *BasesoqlListener) ExitWhereField(ctx *WhereFieldContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BasesoqlListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BasesoqlListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterOrderClause is called when production orderClause is entered.
func (s *BasesoqlListener) EnterOrderClause(ctx *OrderClauseContext) {}

// ExitOrderClause is called when production orderClause is exited.
func (s *BasesoqlListener) ExitOrderClause(ctx *OrderClauseContext) {}

// EnterSoqlValue is called when production soqlValue is entered.
func (s *BasesoqlListener) EnterSoqlValue(ctx *SoqlValueContext) {}

// ExitSoqlValue is called when production soqlValue is exited.
func (s *BasesoqlListener) ExitSoqlValue(ctx *SoqlValueContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BasesoqlListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BasesoqlListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterSoqlFilteringExpression is called when production soqlFilteringExpression is entered.
func (s *BasesoqlListener) EnterSoqlFilteringExpression(ctx *SoqlFilteringExpressionContext) {}

// ExitSoqlFilteringExpression is called when production soqlFilteringExpression is exited.
func (s *BasesoqlListener) ExitSoqlFilteringExpression(ctx *SoqlFilteringExpressionContext) {}

// EnterGroupClause is called when production groupClause is entered.
func (s *BasesoqlListener) EnterGroupClause(ctx *GroupClauseContext) {}

// ExitGroupClause is called when production groupClause is exited.
func (s *BasesoqlListener) ExitGroupClause(ctx *GroupClauseContext) {}

// EnterHavingConditionExpression is called when production havingConditionExpression is entered.
func (s *BasesoqlListener) EnterHavingConditionExpression(ctx *HavingConditionExpressionContext) {}

// ExitHavingConditionExpression is called when production havingConditionExpression is exited.
func (s *BasesoqlListener) ExitHavingConditionExpression(ctx *HavingConditionExpressionContext) {}

// EnterOffsetClause is called when production offsetClause is entered.
func (s *BasesoqlListener) EnterOffsetClause(ctx *OffsetClauseContext) {}

// ExitOffsetClause is called when production offsetClause is exited.
func (s *BasesoqlListener) ExitOffsetClause(ctx *OffsetClauseContext) {}

// EnterViewClause is called when production viewClause is entered.
func (s *BasesoqlListener) EnterViewClause(ctx *ViewClauseContext) {}

// ExitViewClause is called when production viewClause is exited.
func (s *BasesoqlListener) ExitViewClause(ctx *ViewClauseContext) {}

// EnterSoslLiteral is called when production soslLiteral is entered.
func (s *BasesoqlListener) EnterSoslLiteral(ctx *SoslLiteralContext) {}

// ExitSoslLiteral is called when production soslLiteral is exited.
func (s *BasesoqlListener) ExitSoslLiteral(ctx *SoslLiteralContext) {}

// EnterSoslQuery is called when production soslQuery is entered.
func (s *BasesoqlListener) EnterSoslQuery(ctx *SoslQueryContext) {}

// ExitSoslQuery is called when production soslQuery is exited.
func (s *BasesoqlListener) ExitSoslQuery(ctx *SoslQueryContext) {}

// EnterSoslReturningObject is called when production soslReturningObject is entered.
func (s *BasesoqlListener) EnterSoslReturningObject(ctx *SoslReturningObjectContext) {}

// ExitSoslReturningObject is called when production soslReturningObject is exited.
func (s *BasesoqlListener) ExitSoslReturningObject(ctx *SoslReturningObjectContext) {}

// EnterApexIdentifier is called when production apexIdentifier is entered.
func (s *BasesoqlListener) EnterApexIdentifier(ctx *ApexIdentifierContext) {}

// ExitApexIdentifier is called when production apexIdentifier is exited.
func (s *BasesoqlListener) ExitApexIdentifier(ctx *ApexIdentifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasesoqlListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasesoqlListener) ExitLiteral(ctx *LiteralContext) {}
