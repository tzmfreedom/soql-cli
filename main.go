package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Songmu/prompter"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chzyer/readline"
	"github.com/k0kubun/pp"
	"github.com/olekukonko/tablewriter"
	"github.com/tzmfreedom/go-soapforce"
	"github.com/tzmfreedom/soql-cli/parser"
	"github.com/xwb1989/sqlparser"
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
	lexer := parser.NewdmlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewdmlParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Statement()
	return tree.Accept(&StatementBuilder{})
}

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
	where := ctx.WhereClause().GetText()
	return &Statement{
		Type:    "UPDATE",
		Sobject: sobject,
		Values:  values,
		Where:   where,
	}
}

func (v *StatementBuilder) VisitDeleteStatement(ctx *parser.DeleteStatementContext) interface{} {
	sobject := ctx.Sobject().Accept(v).(string)
	where := ctx.WhereClause().GetText()
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
