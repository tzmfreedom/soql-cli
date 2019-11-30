package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/chzyer/readline"
	"github.com/tzmfreedom/go-soapforce"
	"github.com/tzmfreedom/soql-cli/parser"
)

var (
	Version string
)

type option struct {
	Username string
	Password string
	Endpoint string
	Execute  string
	DryRun   bool
}

func main() {
	opt := getOption()
	client := soapforce.NewClient()
	client.SetLoginUrl(opt.Endpoint)
	client.Login(opt.Username, opt.Password)

	var evaluator Evaluator
	if opt.DryRun {
		evaluator = &DryRunner{Client: client}
	} else {
		evaluator = &DefaultEvaluator{Client: client}
	}

	if opt.Execute != "" {
		err := eval(opt.Execute, evaluator)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		return
	}

	repl(evaluator)
}

func repl(evaluator Evaluator) {
	l, _ := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m>>\033[0m ",
		HistoryFile:     "/tmp/land.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})

	for {
		line, err := l.Readline()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		line = strings.TrimSpace(line)
		if line == "exit" || line == "quit" {
			break
		}
		if line == `.schema` {
			// show create table
		}
		if line == `.table` {
			// show database
		}
		if line == "" {
			continue
		}
		err = eval(line, evaluator)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}

func getOption() *option {
	flg := flag.NewFlagSet("soql-cli", flag.ExitOnError)
	username := flg.String("u", "", "username")
	endpoint := flg.String("h", "", "hostname")
	execute := flg.String("e", "", "execute")
	dryrun := flg.Bool("d", false, "dryrun")
	version := flg.Bool("v", false, "version")
	flg.Usage = func() {
		fmt.Printf(`NAME:
   soql - Salesforce Object Query Language CLI
USAGE:
   soql [options]
VERSION:
   %s
OPTIONS:
   -u  username
   -h  hostname (e.g. test.salesforce.com)
   -e  execute command from option
   -d  dryrun (no write, read only mode)
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
		Execute:  *execute,
		DryRun:   *dryrun,
	}
}

func eval(line string, e Evaluator) error {
	if matched, _ := regexp.MatchString(`(?i)^SELECT\s.*\sFROM\s.*`, line); matched {
		e.Select(line)
		return nil
	}
	stmt := parser.ParseString(line)
	switch stmt.Type {
	case "INSERT":
		err := e.Insert(stmt)
		if err != nil {
			return err
		}
	case "UPDATE":
		err := e.Update(stmt)
		if err != nil {
			return err
		}
	case "DELETE":
		err := e.Delete(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
