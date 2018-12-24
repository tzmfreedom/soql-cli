// Code generated from ./parser/dml.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // dml

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by dmlParser.
type dmlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by dmlParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by dmlParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by dmlParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by dmlParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by dmlParser#sobject.
	VisitSobject(ctx *SobjectContext) interface{}

	// Visit a parse tree produced by dmlParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by dmlParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
