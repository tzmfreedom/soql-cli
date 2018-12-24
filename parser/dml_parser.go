// Code generated from ./parser/dml.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // dml

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 24, 80, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 5, 2, 20, 10, 2, 3, 3, 3, 3, 5, 3, 24, 10, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 31, 10, 3, 12, 3, 14, 3, 34, 11, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 42, 10, 3, 12, 3, 14, 3, 45,
	11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 7, 4, 60, 10, 4, 12, 4, 14, 4, 63, 11, 4, 3, 4, 5, 4, 66, 10,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 72, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 3, 4, 2, 17, 19, 21,
	22, 2, 80, 2, 19, 3, 2, 2, 2, 4, 21, 3, 2, 2, 2, 6, 48, 3, 2, 2, 2, 8,
	67, 3, 2, 2, 2, 10, 73, 3, 2, 2, 2, 12, 75, 3, 2, 2, 2, 14, 77, 3, 2, 2,
	2, 16, 20, 5, 4, 3, 2, 17, 20, 5, 6, 4, 2, 18, 20, 5, 8, 5, 2, 19, 16,
	3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 18, 3, 2, 2, 2, 20, 3, 3, 2, 2, 2,
	21, 23, 7, 9, 2, 2, 22, 24, 7, 10, 2, 2, 23, 22, 3, 2, 2, 2, 23, 24, 3,
	2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 26, 5, 10, 6, 2, 26, 27, 7, 3, 2, 2, 27,
	32, 5, 12, 7, 2, 28, 29, 7, 4, 2, 2, 29, 31, 5, 12, 7, 2, 30, 28, 3, 2,
	2, 2, 31, 34, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 35,
	3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 36, 7, 5, 2, 2, 36, 37, 7, 15, 2, 2,
	37, 38, 7, 3, 2, 2, 38, 43, 5, 14, 8, 2, 39, 40, 7, 4, 2, 2, 40, 42, 5,
	14, 8, 2, 41, 39, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43,
	44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 47, 7, 5, 2,
	2, 47, 5, 3, 2, 2, 2, 48, 49, 7, 11, 2, 2, 49, 50, 5, 10, 6, 2, 50, 51,
	7, 16, 2, 2, 51, 52, 5, 12, 7, 2, 52, 53, 7, 6, 2, 2, 53, 61, 5, 14, 8,
	2, 54, 55, 7, 4, 2, 2, 55, 56, 5, 12, 7, 2, 56, 57, 7, 6, 2, 2, 57, 58,
	5, 14, 8, 2, 58, 60, 3, 2, 2, 2, 59, 54, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2,
	61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 65, 3, 2, 2, 2, 63, 61, 3,
	2, 2, 2, 64, 66, 7, 24, 2, 2, 65, 64, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66,
	7, 3, 2, 2, 2, 67, 68, 7, 12, 2, 2, 68, 69, 7, 8, 2, 2, 69, 71, 5, 10,
	6, 2, 70, 72, 7, 24, 2, 2, 71, 70, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72,
	9, 3, 2, 2, 2, 73, 74, 7, 20, 2, 2, 74, 11, 3, 2, 2, 2, 75, 76, 7, 20,
	2, 2, 76, 13, 3, 2, 2, 2, 77, 78, 9, 2, 2, 2, 78, 15, 3, 2, 2, 2, 9, 19,
	23, 32, 43, 61, 65, 71,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "','", "')'", "'='", "'''",
}
var symbolicNames = []string{
	"", "", "", "", "", "QUOTE", "FROM", "INSERT", "INTO", "UPDATE", "DELETE",
	"WHERE", "ON", "VALUES", "SET", "StringLiteral", "BooleanLiteral", "NullLiteral",
	"Identifier", "IntegerLiteral", "FloatingPointLiteral", "WS", "WhereClause",
}

var ruleNames = []string{
	"statement", "insertStatement", "updateStatement", "deleteStatement", "sobject",
	"field", "literal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type dmlParser struct {
	*antlr.BaseParser
}

func NewdmlParser(input antlr.TokenStream) *dmlParser {
	this := new(dmlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "dml.g4"

	return this
}

// dmlParser tokens.
const (
	dmlParserEOF                  = antlr.TokenEOF
	dmlParserT__0                 = 1
	dmlParserT__1                 = 2
	dmlParserT__2                 = 3
	dmlParserT__3                 = 4
	dmlParserQUOTE                = 5
	dmlParserFROM                 = 6
	dmlParserINSERT               = 7
	dmlParserINTO                 = 8
	dmlParserUPDATE               = 9
	dmlParserDELETE               = 10
	dmlParserWHERE                = 11
	dmlParserON                   = 12
	dmlParserVALUES               = 13
	dmlParserSET                  = 14
	dmlParserStringLiteral        = 15
	dmlParserBooleanLiteral       = 16
	dmlParserNullLiteral          = 17
	dmlParserIdentifier           = 18
	dmlParserIntegerLiteral       = 19
	dmlParserFloatingPointLiteral = 20
	dmlParserWS                   = 21
	dmlParserWhereClause          = 22
)

// dmlParser rules.
const (
	dmlParserRULE_statement       = 0
	dmlParserRULE_insertStatement = 1
	dmlParserRULE_updateStatement = 2
	dmlParserRULE_deleteStatement = 3
	dmlParserRULE_sobject         = 4
	dmlParserRULE_field           = 5
	dmlParserRULE_literal         = 6
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dmlParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dmlParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) InsertStatement() IInsertStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsertStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsertStatementContext)
}

func (s *StatementContext) UpdateStatement() IUpdateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateStatementContext)
}

func (s *StatementContext) DeleteStatement() IDeleteStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeleteStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dmlVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dmlParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, dmlParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case dmlParserINSERT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.InsertStatement()
		}

	case dmlParserUPDATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(15)
			p.UpdateStatement()
		}

	case dmlParserDELETE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(16)
			p.DeleteStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInsertStatementContext is an interface to support dynamic dispatch.
type IInsertStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsertStatementContext differentiates from other interfaces.
	IsInsertStatementContext()
}

type InsertStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertStatementContext() *InsertStatementContext {
	var p = new(InsertStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dmlParserRULE_insertStatement
	return p
}

func (*InsertStatementContext) IsInsertStatementContext() {}

func NewInsertStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertStatementContext {
	var p = new(InsertStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dmlParserRULE_insertStatement

	return p
}

func (s *InsertStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertStatementContext) INSERT() antlr.TerminalNode {
	return s.GetToken(dmlParserINSERT, 0)
}

func (s *InsertStatementContext) Sobject() ISobjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISobjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISobjectContext)
}

func (s *InsertStatementContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *InsertStatementContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *InsertStatementContext) VALUES() antlr.TerminalNode {
	return s.GetToken(dmlParserVALUES, 0)
}

func (s *InsertStatementContext) AllLiteral() []ILiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralContext)(nil)).Elem())
	var tst = make([]ILiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralContext)
		}
	}

	return tst
}

func (s *InsertStatementContext) Literal(i int) ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *InsertStatementContext) INTO() antlr.TerminalNode {
	return s.GetToken(dmlParserINTO, 0)
}

func (s *InsertStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.EnterInsertStatement(s)
	}
}

func (s *InsertStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.ExitInsertStatement(s)
	}
}

func (s *InsertStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dmlVisitor:
		return t.VisitInsertStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dmlParser) InsertStatement() (localctx IInsertStatementContext) {
	localctx = NewInsertStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, dmlParserRULE_insertStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Match(dmlParserINSERT)
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == dmlParserINTO {
		{
			p.SetState(20)
			p.Match(dmlParserINTO)
		}

	}
	{
		p.SetState(23)
		p.Sobject()
	}
	{
		p.SetState(24)
		p.Match(dmlParserT__0)
	}
	{
		p.SetState(25)
		p.Field()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == dmlParserT__1 {
		{
			p.SetState(26)
			p.Match(dmlParserT__1)
		}
		{
			p.SetState(27)
			p.Field()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Match(dmlParserT__2)
	}
	{
		p.SetState(34)
		p.Match(dmlParserVALUES)
	}
	{
		p.SetState(35)
		p.Match(dmlParserT__0)
	}
	{
		p.SetState(36)
		p.Literal()
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == dmlParserT__1 {
		{
			p.SetState(37)
			p.Match(dmlParserT__1)
		}
		{
			p.SetState(38)
			p.Literal()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
		p.Match(dmlParserT__2)
	}

	return localctx
}

// IUpdateStatementContext is an interface to support dynamic dispatch.
type IUpdateStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdateStatementContext differentiates from other interfaces.
	IsUpdateStatementContext()
}

type UpdateStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateStatementContext() *UpdateStatementContext {
	var p = new(UpdateStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dmlParserRULE_updateStatement
	return p
}

func (*UpdateStatementContext) IsUpdateStatementContext() {}

func NewUpdateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateStatementContext {
	var p = new(UpdateStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dmlParserRULE_updateStatement

	return p
}

func (s *UpdateStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateStatementContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(dmlParserUPDATE, 0)
}

func (s *UpdateStatementContext) Sobject() ISobjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISobjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISobjectContext)
}

func (s *UpdateStatementContext) SET() antlr.TerminalNode {
	return s.GetToken(dmlParserSET, 0)
}

func (s *UpdateStatementContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *UpdateStatementContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *UpdateStatementContext) AllLiteral() []ILiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralContext)(nil)).Elem())
	var tst = make([]ILiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralContext)
		}
	}

	return tst
}

func (s *UpdateStatementContext) Literal(i int) ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *UpdateStatementContext) WhereClause() antlr.TerminalNode {
	return s.GetToken(dmlParserWhereClause, 0)
}

func (s *UpdateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.EnterUpdateStatement(s)
	}
}

func (s *UpdateStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.ExitUpdateStatement(s)
	}
}

func (s *UpdateStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dmlVisitor:
		return t.VisitUpdateStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dmlParser) UpdateStatement() (localctx IUpdateStatementContext) {
	localctx = NewUpdateStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, dmlParserRULE_updateStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(dmlParserUPDATE)
	}
	{
		p.SetState(47)
		p.Sobject()
	}
	{
		p.SetState(48)
		p.Match(dmlParserSET)
	}
	{
		p.SetState(49)
		p.Field()
	}
	{
		p.SetState(50)
		p.Match(dmlParserT__3)
	}
	{
		p.SetState(51)
		p.Literal()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == dmlParserT__1 {
		{
			p.SetState(52)
			p.Match(dmlParserT__1)
		}
		{
			p.SetState(53)
			p.Field()
		}
		{
			p.SetState(54)
			p.Match(dmlParserT__3)
		}
		{
			p.SetState(55)
			p.Literal()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == dmlParserWhereClause {
		{
			p.SetState(62)
			p.Match(dmlParserWhereClause)
		}

	}

	return localctx
}

// IDeleteStatementContext is an interface to support dynamic dispatch.
type IDeleteStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteStatementContext differentiates from other interfaces.
	IsDeleteStatementContext()
}

type DeleteStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteStatementContext() *DeleteStatementContext {
	var p = new(DeleteStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dmlParserRULE_deleteStatement
	return p
}

func (*DeleteStatementContext) IsDeleteStatementContext() {}

func NewDeleteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStatementContext {
	var p = new(DeleteStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dmlParserRULE_deleteStatement

	return p
}

func (s *DeleteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStatementContext) DELETE() antlr.TerminalNode {
	return s.GetToken(dmlParserDELETE, 0)
}

func (s *DeleteStatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(dmlParserFROM, 0)
}

func (s *DeleteStatementContext) Sobject() ISobjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISobjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISobjectContext)
}

func (s *DeleteStatementContext) WhereClause() antlr.TerminalNode {
	return s.GetToken(dmlParserWhereClause, 0)
}

func (s *DeleteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.EnterDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.ExitDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dmlVisitor:
		return t.VisitDeleteStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dmlParser) DeleteStatement() (localctx IDeleteStatementContext) {
	localctx = NewDeleteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, dmlParserRULE_deleteStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(dmlParserDELETE)
	}
	{
		p.SetState(66)
		p.Match(dmlParserFROM)
	}
	{
		p.SetState(67)
		p.Sobject()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == dmlParserWhereClause {
		{
			p.SetState(68)
			p.Match(dmlParserWhereClause)
		}

	}

	return localctx
}

// ISobjectContext is an interface to support dynamic dispatch.
type ISobjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSobjectContext differentiates from other interfaces.
	IsSobjectContext()
}

type SobjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySobjectContext() *SobjectContext {
	var p = new(SobjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dmlParserRULE_sobject
	return p
}

func (*SobjectContext) IsSobjectContext() {}

func NewSobjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SobjectContext {
	var p = new(SobjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dmlParserRULE_sobject

	return p
}

func (s *SobjectContext) GetParser() antlr.Parser { return s.parser }

func (s *SobjectContext) Identifier() antlr.TerminalNode {
	return s.GetToken(dmlParserIdentifier, 0)
}

func (s *SobjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SobjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SobjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.EnterSobject(s)
	}
}

func (s *SobjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.ExitSobject(s)
	}
}

func (s *SobjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dmlVisitor:
		return t.VisitSobject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dmlParser) Sobject() (localctx ISobjectContext) {
	localctx = NewSobjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, dmlParserRULE_sobject)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(dmlParserIdentifier)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dmlParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dmlParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Identifier() antlr.TerminalNode {
	return s.GetToken(dmlParserIdentifier, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dmlVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dmlParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, dmlParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(dmlParserIdentifier)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dmlParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dmlParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(dmlParserIntegerLiteral, 0)
}

func (s *LiteralContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(dmlParserFloatingPointLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(dmlParserStringLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(dmlParserBooleanLiteral, 0)
}

func (s *LiteralContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(dmlParserNullLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dmlListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dmlVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dmlParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, dmlParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<dmlParserStringLiteral)|(1<<dmlParserBooleanLiteral)|(1<<dmlParserNullLiteral)|(1<<dmlParserIntegerLiteral)|(1<<dmlParserFloatingPointLiteral))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
