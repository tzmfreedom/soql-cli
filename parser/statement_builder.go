package parser

type Statement struct {
	Type    string
	Sobject string
	Values  map[string]string
	Where   string
}

type StatementBuilder struct {
	*BasedmlVisitor
}

func (v *StatementBuilder) VisitStatement(ctx *StatementContext) interface{} {
	if stmt := ctx.InsertStatement(); stmt != nil {
		return stmt.Accept(v)
	}
	if stmt := ctx.UpdateStatement(); stmt != nil {
		return stmt.Accept(v)
	}
	return ctx.DeleteStatement().Accept(v)
}

func (v *StatementBuilder) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
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

func (v *StatementBuilder) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
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

func (v *StatementBuilder) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
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

func (v *StatementBuilder) VisitSobject(ctx *SobjectContext) interface{} {
	return ctx.Identifier().GetText()
}

func (v *StatementBuilder) VisitField(ctx *FieldContext) interface{} {
	return ctx.Identifier().GetText()
}

func (v *StatementBuilder) VisitLiteral(ctx *LiteralContext) interface{} {
	if s := ctx.StringLiteral(); s != nil {
		str := s.GetText()
		return str[1 : len(str)-1]
	}
	return ctx.GetText()
}
