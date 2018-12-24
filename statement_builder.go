package main

import "github.com/tzmfreedom/soql-cli/parser"

type Statement struct {
	Type    string
	Sobject string
	Values  map[string]string
	Where   string
}

type StatementBuilder struct {
	*parser.BasedmlVisitor
}

func (v *StatementBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
	if stmt := ctx.InsertStatement(); stmt != nil {
		return stmt.Accept(v)
	}
	if stmt := ctx.UpdateStatement(); stmt != nil {
		return stmt.Accept(v)
	}
	return ctx.DeleteStatement().Accept(v)
}

func (v *StatementBuilder) VisitInsertStatement(ctx *parser.InsertStatementContext) interface{} {
	values := map[string]string{}
	for i, f := range ctx.AllField() {
		key := f.Accept(v).(string)
		val := ctx.Literal(i).Accept(v).(string)
		values[key] = val
	}
	sobject := ctx.Sobject().Accept(v).(string)
	return &Statement{
		Type:    "INSERT",
		Sobject: sobject,
		Values:  values,
	}
}

func (v *StatementBuilder) VisitUpdateStatement(ctx *parser.UpdateStatementContext) interface{} {
	values := map[string]string{}
	for i, f := range ctx.AllField() {
		key := f.Accept(v).(string)
		val := ctx.Literal(i).Accept(v).(string)
		values[key] = val
	}
	sobject := ctx.Sobject().Accept(v).(string)
	var where string
	if w := ctx.WhereClause(); w != nil {
		where = ctx.WhereClause().GetText()
	}
	return &Statement{
		Type:    "UPDATE",
		Sobject: sobject,
		Values:  values,
		Where:   where,
	}
}

func (v *StatementBuilder) VisitDeleteStatement(ctx *parser.DeleteStatementContext) interface{} {
	sobject := ctx.Sobject().Accept(v).(string)
	var where string
	if w := ctx.WhereClause(); w != nil {
		where = ctx.WhereClause().GetText()
	}
	return &Statement{
		Type:    "DELETE",
		Sobject: sobject,
		Where:   where,
	}
}

func (v *StatementBuilder) VisitSobject(ctx *parser.SobjectContext) interface{} {
	return ctx.Identifier().GetText()
}

func (v *StatementBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	return ctx.Identifier().GetText()
}

func (v *StatementBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if s := ctx.StringLiteral(); s != nil {
		str := s.GetText()
		return str[1 : len(str)-1]
	}
	return ctx.GetText()
}
