// Code generated from ./parser/dml.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // dml

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasedmlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasedmlVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedmlVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedmlVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedmlVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedmlVisitor) VisitSobject(ctx *SobjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedmlVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedmlVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
