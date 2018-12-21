package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Songmu/prompter"
	"github.com/chzyer/readline"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/go-soapforce"
	"github.com/tzmfreedom/soql-cli/parser"
	"github.com/xwb1989/sqlparser"
	"github.com/olekukonko/tablewriter"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var (
	Version string
)

type option struct {
	Username string
	Password string
	Endpoint string
}

func main() {
	opt := getOption()
	client := soapforce.NewClient()
	client.SetLoginUrl(opt.Endpoint)
	client.Login(opt.Username, opt.Password)

	l, _ := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m>>\033[0m ",
		HistoryFile:     "/tmp/land.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})

	for {
		line, err := l.Readline()
		if err != nil {
			panic(err)
		}
		stmt, err := sqlparser.Parse(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		pp.Println(stmt.(*sqlparser.Select))

		os.Exit(0)
		switch stmt := stmt.(type) {
		case *sqlparser.Select:
			r, err := client.Query(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				continue
			}
			pp.Println(stmt)
			pp.Println(r)
			fieldNames := []string{}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader(fieldNames)

			for _, record := range r.Records {
				columns := []string{}
				for _, fieldName := range fieldNames {
					if record.Id != "" {
						columns = append(columns, record.Id)
					} else {
						columns = append(columns, record.Fields[fieldName].(string))
					}
				}
				table.Append(columns)
			}
		case *sqlparser.Insert:
		case *sqlparser.Update:
		case *sqlparser.Delete:
		}
	}
}

func getOption() *option {
	flg := flag.NewFlagSet("lui", flag.ExitOnError)
	username := flg.String("u", "", "username")
	endpoint := flg.String("e", "", "endpoint")
	version := flg.Bool("v", false, "version")
	flg.Usage = func() {
		fmt.Printf(`NAME:
   lui - Lightining plaform terminal UI
USAGE:
   lui [options]
VERSION:
   %s
OPTIONS:
   -u  username
   -p  username
   -e  endpoint (e.g. test.salesforce.com)
   -v  print the version
`, Version)
		os.Exit(0)
	}
	err := flg.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	if *version {
		fmt.Printf("%s\n", Version)
		os.Exit(0)
		return nil
	}

	if *username == "" && os.Getenv("SALESFORCE_USERNAME") != "" {
		*username = os.Getenv("SALESFORCE_USERNAME")
	}
	if *username == "" {
		*username = prompter.Prompt("Enter username", "")
	}
	if *endpoint == "" && os.Getenv("SALESFORCE_ENDPOINT") != "" {
		*endpoint = os.Getenv("SALESFORCE_ENDPOINT")
	}
	if *endpoint == "" {
		*endpoint = prompter.Prompt("Enter endpoint (e.g. test.salesforce.com)", "login.salesforce.com")
	}

	password := os.Getenv("SALESFORCE_PASSWORD")
	if password == "" {
		password = prompter.Password("Enter password")
	}

	return &option{
		Username: *username,
		Password: password,
		Endpoint: *endpoint,
	}
}

func ParseFile(f string) (interface{}, error) {
	input, err := antlr.NewFileStream(f)
	if err != nil {
		return nil, err
	}
	return parse(input), nil
}

func ParseString(data string) interface{} {
	input := antlr.NewInputStream(data)
	return parse(input)
}

func parse(input antlr.CharStream) interface{} {
	lexer := parser.NewsoqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewsoqlParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Query()
	return tree.Accept(&SoqlBuilder{})
}

type Select struct {
	Fields []*Field
	Where []*Where
}

type Field struct {
	Name []string
}

type Where struct{}

type SoqlBuilder struct {
	*antlr.BaseParseTreeVisitor
}

func (v *SoqlBuilder) VisitQuery(ctx *parser.QueryContext) interface{} {
	selectClause := ctx.SelectClause().Accept(v)
	return selectClause
}

func (v *SoqlBuilder) VisitSelectClause(ctx *parser.SelectClauseContext) interface{} {
	return ctx.FieldList().Accept(v)
}

func (v *SoqlBuilder) VisitFieldList(ctx *parser.FieldListContext) interface{} {
	selectFields := ctx.AllSelectField()
	fields := make([]*Field, len(selectFields))
	for i, f := range selectFields {
		fields[i] = f.Accept(v).(*Field)
	}
	return fields
}

func (v *SoqlBuilder) VisitSelectField(ctx *parser.SelectFieldContext) interface{} {
	if f := ctx.SoqlField(); f != nil {
		return f.Accept(v)
	}
	if f := ctx.Subquery(); f != nil {
		return f.Accept(v)
	}
	return nil
}

func (v *SoqlBuilder) VisitFromClause(ctx *parser.FromClauseContext) interface{} {
	return ctx.ApexIdentifier().GetText()
}

func (v *SoqlBuilder) VisitFilterScope(ctx *parser.FilterScopeContext) interface{} {
	return nil
}

func (v *SoqlBuilder) VisitSoqlFieldReference(ctx *parser.SoqlFieldReferenceContext) interface{} {
	identifiers := ctx.AllApexIdentifier()
	names := make([]string, len(identifiers))
	for i, ident := range identifiers {
		names[i] = ident.GetText()
	}
	return names
}

func (v *SoqlBuilder) VisitSoqlFunctionCall(ctx *parser.SoqlFunctionCallContext) interface{} {
	funcName := ctx.ApexIdentifier().GetText()
	fields := ctx.AllSoqlField()
	args := make([]*Field, len(fields))
	for i, f := range fields {
		args[i] = f.Accept(v).(*Field)
	}
	return nil
}

func (v *SoqlBuilder) VisitSubquery(ctx *parser.SubqueryContext) interface{} {
	return ctx.Query().Accept(v)
}

func (v *SoqlBuilder) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitWhereFields(ctx *parser.WhereFieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitWhereField(ctx *parser.WhereFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitLimitClause(ctx *parser.LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitOrderClause(ctx *parser.OrderClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitSoqlValue(ctx *parser.SoqlValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitWithClause(ctx *parser.WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitSoqlFilteringExpression(ctx *parser.SoqlFilteringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitGroupClause(ctx *parser.GroupClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitHavingConditionExpression(ctx *parser.HavingConditionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitOffsetClause(ctx *parser.OffsetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitViewClause(ctx *parser.ViewClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitSoslLiteral(ctx *parser.SoslLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitSoslQuery(ctx *parser.SoslQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitSoslReturningObject(ctx *parser.SoslReturningObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitApexIdentifier(ctx *parser.ApexIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
