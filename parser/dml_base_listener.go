// Code generated from ./parser/dml.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // dml

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasedmlListener is a complete listener for a parse tree produced by dmlParser.
type BasedmlListener struct{}

var _ dmlListener = &BasedmlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasedmlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasedmlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasedmlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasedmlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasedmlListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasedmlListener) ExitStatement(ctx *StatementContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BasedmlListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BasedmlListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BasedmlListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BasedmlListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BasedmlListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BasedmlListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterSobject is called when production sobject is entered.
func (s *BasedmlListener) EnterSobject(ctx *SobjectContext) {}

// ExitSobject is called when production sobject is exited.
func (s *BasedmlListener) ExitSobject(ctx *SobjectContext) {}

// EnterField is called when production field is entered.
func (s *BasedmlListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BasedmlListener) ExitField(ctx *FieldContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasedmlListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasedmlListener) ExitLiteral(ctx *LiteralContext) {}
