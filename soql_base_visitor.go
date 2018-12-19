// Code generated from soql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // soql

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasesoqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasesoqlVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitSelectClause(ctx *SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitFieldList(ctx *FieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitSelectField(ctx *SelectFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitFilterScope(ctx *FilterScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitSoqlFieldReference(ctx *SoqlFieldReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitSoqlFunctionCall(ctx *SoqlFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitWhereFields(ctx *WhereFieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitWhereField(ctx *WhereFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitOrderClause(ctx *OrderClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitSoqlValue(ctx *SoqlValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitSoqlFilteringExpression(ctx *SoqlFilteringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitGroupClause(ctx *GroupClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitHavingConditionExpression(ctx *HavingConditionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitOffsetClause(ctx *OffsetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitViewClause(ctx *ViewClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitSoslLiteral(ctx *SoslLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitSoslQuery(ctx *SoslQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitSoslReturningObject(ctx *SoslReturningObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitApexIdentifier(ctx *ApexIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesoqlVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
