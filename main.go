package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Songmu/prompter"
	"github.com/chzyer/readline"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/go-soapforce"
	"github.com/xwb1989/sqlparser"
	"github.com/olekukonko/tablewriter"
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
