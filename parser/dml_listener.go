// Code generated from ./parser/dml.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // dml

import "github.com/antlr/antlr4/runtime/Go/antlr"

// dmlListener is a complete listener for a parse tree produced by dmlParser.
type dmlListener interface {
	antlr.ParseTreeListener

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterSobject is called when entering the sobject production.
	EnterSobject(c *SobjectContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitSobject is called when exiting the sobject production.
	ExitSobject(c *SobjectContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
